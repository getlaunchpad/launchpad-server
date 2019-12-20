package jwt

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-chi/jwtauth"
)

type Token struct {
	Claims    Claims
	tokenAuth *jwtauth.JWTAuth
}

type Claims struct {
	jwt.StandardClaims
	UserID uint
	Role   string
}

// Creates new jwt with valid jwtauth parcer
func (Token) New() *Token {
	token := &Token{
		tokenAuth: jwtauth.New("HS256", []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil),
	}

	return token
}

func (t *Token) Encode(UserID uint, Role string) string {
	claims := &Claims{
		UserID: UserID,
		Role:   Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "getlaunchpad.dev",
		},
	}

	_, tokenString, err := t.tokenAuth.Encode(claims)
	if err != nil {
		log.Panic(err)
		return ""
	}

	return tokenString
}
