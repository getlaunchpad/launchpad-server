package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lucasstettner/launchpad-server/app/constants"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Check if we get redirected when logging in
func TestGoogleOauthLogin(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/auth/google/login", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusTemporaryRedirect, response.Code)
}

// Ensure that we get denied access if a state is not present
func TestGoogleOauthCallbackInvalidState(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/auth/google/callback", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	r := Response{}
	if err := json.Unmarshal(response.Body.Bytes(), &r); err != nil {
		t.Errorf("%s\n", constants.DecodeRequestBodyErr)
	}

	if r.Success != false && r.Message != "Invalid oauth google state" {
		t.Error("Unexpected api response")
	}
}
