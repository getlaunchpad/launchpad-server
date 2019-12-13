package app

import (
	"net/http"

	"github.com/lucasstettner/launchpad-server/app/handlers"
)

func (s *Server) initializeRoutes() {
	// These routes are used for kubernetes
	s.Router.Get("/health", handlers.HealthHandler)
	s.Router.Get("/readiness", handlers.ReadinessHandler)

	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
