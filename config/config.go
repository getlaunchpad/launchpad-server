package config

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/fsnotify/fsnotify"
	"github.com/lucasstettner/launchpad-server/app/models"
	"github.com/spf13/viper"

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
	config := &Config{}
	constants, err := initViper()
	config.Constants = constants
	if err != nil {
		log.Fatal(err)
		return &Config{}
	}

	// Set Google Oauth Info
	config.Constants.GConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  viper.GetString("GoogleRedirectUrl"),
	}

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_PASSWORD"))

	config.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		log.Fatal(err)
		return &Config{}
	}

	// turn this off in prod
	config.DB.LogMode(true)

	// Create necessary types (such as roles for user) before migration
	// config.DB.Exec("CREATE TYPE role AS ENUM ('member','pro');")

	config.DB.AutoMigrate(&models.User{})

	return config
}

func initViper() (Constants, error) {
	viper.SetConfigName("launchpad.config") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")                // Search the root directory for the configuration file
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {                         // Handle errors reading the config file
		log.Fatal(err)
		return Constants{}, err
	}
	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.SetDefault("PORT", "8080")
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
