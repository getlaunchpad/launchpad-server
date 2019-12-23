package config

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/lucasstettner/launchpad-server/app/models"

	"github.com/jinzhu/gorm"
)

type Constants struct {
	Version string
	GConfig *oauth2.Config
}

type Config struct {
	Constants Constants
	DB        *gorm.DB
}

func New() *Config {
	var err error

	config := &Config{}

	// Set Google Oauth Info
	config.Constants.GConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/v1/auth/google/callback",
	}

	config.DB, err = gorm.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
		return &Config{}
	}

	// turn this off in prod
	config.DB.LogMode(true)

	// Create necessary types (such as roles for user) before migration
	config.DB.Exec("CREATE TYPE role AS ENUM ('member','pro');")

	config.DB.AutoMigrate(&models.User{})

	return config
}
