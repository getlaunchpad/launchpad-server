package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/lucasstettner/launchpad-server/config"
)

// This tests the main function of the app, it is necessary that this passes
// Therefore we use the testing.Main package
func TestMain(m *testing.M) {
	if err := godotenv.Load(os.ExpandEnv("./.env")); err != nil {
		log.Printf("Error getting env, continuing in production mode %v\n", err)
	}

	// Integration Test DB Connection and Viper configuration
	config.New()

	os.Exit(m.Run())
}
