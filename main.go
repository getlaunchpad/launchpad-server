package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lucasstettner/launchpad-server/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		// This currently breaks production, because env vars are handled in circleci
		log.Panic(err)
	} else {
		log.Println("Fetching env variables")
	}

	app.Start()
}
