package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/lucasstettner/launchpad-server/app/utils/jwt"

	"github.com/lucasstettner/launchpad-server/app/models"
	"golang.org/x/oauth2"

	"github.com/lucasstettner/launchpad-server/app/constants"

	"github.com/go-chi/chi"
	"github.com/lucasstettner/launchpad-server/config"
)

const userInfoURI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type Config struct {
	*config.Config
}

func New(c *config.Config) *Config {
	return &Config{c}
}

func (c *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/login", c.oauthGoogleLogin)
	router.Get("/callback", c.oauthGoogleCallback)
	return router
}

type googleAuthResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Redirect user to google oauth flpw
func (c *Config) oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
	   AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	   validate that it matches the the state query parameter on your redirect callback.
	*/
	u := c.Constants.GConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

/*
	We redirect user back even if there is an error to avoid leaving
	their browser on the api
*/
func (c *Config) oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, err := r.Cookie("oauthstate")
	if err != nil || oauthState.Value == "" {
		redirectToFrontend(w, r, "invalid oauth state")
		return
	}

	if r.FormValue("state") != oauthState.Value {
		redirectToFrontend(w, r, "invalid oauth state")
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"), c.Constants.GConfig)
	if err != nil {
		redirectToFrontend(w, r, err.Error())
		return
	}

	guser := googleAuthResponse{}
	if err := json.Unmarshal(data, &guser); err != nil {
		redirectToFrontend(w, r, constants.DecodeRequestBodyErr)
		return
	}

	user := models.User{
		Email:    guser.Email,
		GoogleID: guser.ID,
	}
	if err := user.LoginOrSignup(c.DB); err != nil {
		redirectToFrontend(w, r, "error logging in")
		return
	}

	// Create/Encode jwt token with user id and role
	t := jwt.Token{}.New().Encode(user.ID, user.Role)

	// Add 'token' cookie to request header response
	setAuthCookie(w, t)

	// Redirect to frontend with auth header
	http.Redirect(w, r, "http://localhost:3000/", http.StatusTemporaryRedirect)
}

// Generates state cookie under oauthstate
func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string, GConfig *oauth2.Config) ([]byte, error) {
	// Use code to get token and get user info from Google.
	token, err := GConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(userInfoURI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func setAuthCookie(w http.ResponseWriter, t string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   t,
		Path:    "/",
		Expires: time.Now().Add(5 * time.Minute),
	})
}

func redirectToFrontend(w http.ResponseWriter, r *http.Request, err string) {
	http.Redirect(w, r, "http://localhost:3000/login?error="+url.QueryEscape(err), http.StatusTemporaryRedirect)
}
