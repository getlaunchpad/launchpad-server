package tests

import (
	"testing"

	"github.com/lucasstettner/launchpad-server/app/models"
)

func TestLoginOrSignupSuccess(t *testing.T) {
	user := &models.User{
		Email:    "test@gmail.com",
		GoogleID: "12345",
	}
	if err := user.LoginOrSignup(a.Config.DB); err != nil {
		t.Errorf("Error Login/Signup: %s", err)
	}
}
