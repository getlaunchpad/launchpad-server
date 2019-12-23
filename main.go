package main

import (
	"log"

	"github.com/lucasstettner/launchpad-server/app"

	"github.com/joho/godotenv"
)

func main() {
	a := app.App{}

	if err := godotenv.Load(); err != nil {
		log.Println(err)
	} else {
		log.Println("Fetching env variables")
	}

	a.Start(true)
}
