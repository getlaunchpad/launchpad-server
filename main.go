package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lucasstettner/launchpad-server/app"
)

var server = app.Server{}

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Fetching env variables")
	}

	server.Initialize()
}
