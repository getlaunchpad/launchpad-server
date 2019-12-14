package app

import (
	"fmt"
	"net/http"

	"github.com/lucasstettner/launchpad-server/app/handlers"
)

func (s *Server) initializeRoutes() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello world!")
	})

	// These routes are used for kubernetes
	s.Router.Get("/health", handlers.HealthHandler)
	s.Router.Get("/readiness", handlers.ReadinessHandler)
}
