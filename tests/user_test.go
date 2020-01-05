package tests

import (
	"net/http"
	"testing"

	"github.com/lucasstettner/launchpad-server/app/constants"
	"github.com/lucasstettner/launchpad-server/app/utils/responses"

	"encoding/json"

	"github.com/lucasstettner/launchpad-server/app/models"
	"github.com/lucasstettner/launchpad-server/app/utils/jwt"
)

type successfulResponse struct {
	Data models.User
}

// Test login or signup success
func TestLoginOrSignupSuccess(t *testing.T) {
	user := &models.User{
		Email:    "test@gmail.com",
		GoogleID: "12345",
	}
	if err := user.LoginOrSignup(a.Config.DB); err != nil {
		t.Errorf("Error Login/Signup: %s", err)
	}
}

func TestMeRouteUnauthorized(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/user/me", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestMeRouteUserNotFound(t *testing.T) {
	token := jwt.Token{}.New()
	// 0 will never be a valid user ID
	tokenstr := token.Encode(0, "member")

	req, _ := http.NewRequest("GET", "/v1/user/me", nil)
	req.Header.Add("Authorization", "Bearer "+tokenstr)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	r := responses.Response{}
	if err := json.Unmarshal(response.Body.Bytes(), &r); err != nil {
		t.Errorf("%s\n", constants.DecodeRequestBodyErr)
	}

	if r.Error.Message != "User Not Found" {
		t.Errorf("API returned %s when 'User Not Found' expected", r.Error.Message)
	}
}

func TestMeRouteAuthorized(t *testing.T) {
	token := jwt.Token{}.New()
	tokenstr := token.Encode(1, "member")

	req, _ := http.NewRequest("GET", "/v1/user/me", nil)
	req.Header.Add("Authorization", "Bearer "+tokenstr)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	r := successfulResponse{}
	if err := json.Unmarshal(response.Body.Bytes(), &r); err != nil {
		t.Errorf("%s\n", constants.DecodeRequestBodyErr)
	}

	if r.Data.ID != 1 {
		t.Errorf("Expected ID of 1 but got %v instead", r.Data.ID)
	} else if r.Data.Role != "member" {
		t.Errorf("Expected Role of 'member' but got %v instead", r.Data.Role)
	}
}
