package tests

import (
	"net/http"
	"testing"
)

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

	checkResponseCode(t, http.StatusTemporaryRedirect, response.Code)
}
