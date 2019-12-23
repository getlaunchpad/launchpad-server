package tests

import (
	"net/http"
	"testing"
)

func TestGoogleOauthLogin(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/auth/google/login", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusTemporaryRedirect, response.Code)
}
