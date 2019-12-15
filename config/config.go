package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	OauthConfig *oauth2.Config
	UserInfoUrl string
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

type Config struct {
	Google   GoogleConfig
	DBConfig DBConfig
	Database *gorm.DB
}

func New() *Config {
	var err error
	config := &Config{
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
		DBConfig: DBConfig{
			Dialect:  "postgres",
			Host:     "127.0.0.1",
			Port:     "1234",
			Username: "postgres",
			Password: "lucasthebear",
			Name:     "launchpad",
			Charset:  "utf8",
		},
	}

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DBConfig.Host,
		config.DBConfig.Port,
		config.DBConfig.Username,
		config.DBConfig.Name,
		config.DBConfig.Password)

	config.Database, err = gorm.Open(config.DBConfig.Dialect, DBURL)
	if err != nil {
		log.Fatal(err)
		return &Config{}
	}

	// turn this off in prod
	config.Database.LogMode(true)

	defer config.Database.Close()

	// db.AutoMigrate(&models.User{})

	return config
}
