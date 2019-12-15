package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	OauthConfig *oauth2.Config
	UserInfoUrl string
}

type Config struct {
	Google GoogleConfig
}

func New() *Config {
	return &Config{
		Google: GoogleConfig{
			OauthConfig: &oauth2.Config{
				RedirectURL:  "http://localhost:8080/v1/auth/google/callback",
				ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
				ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
				Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
				Endpoint:     google.Endpoint,
			},
			UserInfoUrl: "https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
		},
	}
}
