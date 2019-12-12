package main

import (
	"github.com/lucasstettner/launchpad-server/app"
)

var server = app.Server{}

func main() {
	server.Initialize()
}
