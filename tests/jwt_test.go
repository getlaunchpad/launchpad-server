package tests

import (
	"testing"
	"time"

	"github.com/lucasstettner/launchpad-server/app/utils/jwt"
)

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
