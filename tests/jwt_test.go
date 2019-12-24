package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/lucasstettner/launchpad-server/app/utils/jwt"
)

func TestAdminUnauthorized(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/admin", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestAdminAuthorized(t *testing.T) {
	token := jwt.Token{}.New()
	tokenstr := token.Encode(2, "member")

	req, _ := http.NewRequest("GET", "/v1/admin", nil)
	req.Header.Add("Authorization", "Bearer "+tokenstr)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestJWTEncoding(t *testing.T) {
	// Encode jwt string with claims
	tokenHandler := jwt.Token{}.New()
	tokenstr := tokenHandler.Encode(21, "member")

	claims, err := tokenHandler.ParseToken(tokenstr)
	if err != nil {
		t.Errorf("Error decrypting jwt token: %v", err)
	}

	if err = claims.Valid(); err != nil {
		t.Errorf("Token is not valid: %v\n", err)
	}

	if claims.UserID != 21 {
		t.Error("UserID does not match")
	}

	if claims.ExpiresAt != time.Now().Add(time.Minute*15).Unix() {
		t.Error("ExpiresAt is not correct")
	}

	if claims.IssuedAt != time.Now().Unix() {
		t.Error("IssuedAt is not correct")
	}
}
