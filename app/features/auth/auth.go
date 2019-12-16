package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lucasstettner/launchpad-server/app/models"

	"github.com/lucasstettner/launchpad-server/app/constants"

	"github.com/go-chi/chi"
	"github.com/lucasstettner/launchpad-server/app/utils/responses"
	"github.com/lucasstettner/launchpad-server/config"
)

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

type GoogleAuthResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Redirect user to google oauth flpw
func (c *Config) oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	config := config.New()

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
	   AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	   validate that it matches the the state query parameter on your redirect callback.
	*/
	u := config.Google.OauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

// Callback from google flow
func (c *Config) oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		responses.Error(w, http.StatusBadRequest, "Invalid oauth google state")
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	guser := GoogleAuthResponse{}
	if err := json.Unmarshal(data, &guser); err != nil {
		responses.Error(w, http.StatusBadRequest, constants.DecodeRequestBodyErr)
		return
	}

	user := models.User{
		Email:    guser.Email,
		GoogleID: guser.ID,
	}
	if err := user.LoginOrSignup(c.DB); err != nil {
		responses.Error(w, http.StatusBadRequest, "Error logging/signing up")
		return
	}

	// Print out user details
	// This is temporary, later down the line we can do a LoginOrSignup
	fmt.Fprintf(w, "UserInfo: %s\n", guser)
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

func getUserDataFromGoogle(code string) ([]byte, error) {
	config := config.New()

	// Use code to get token and get user info from Google.
	token, err := config.Google.OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(config.Google.UserInfoUrl + token.AccessToken)
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
