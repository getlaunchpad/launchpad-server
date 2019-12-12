package app

import (
	"net/http"

	"github.com/lucasstettner/launchpad-server/app/handlers"

	"github.com/go-chi/chi"
)

func (s *Server) initializeRoutes() {
	// These routes are used for kubernetes
	s.Router.Get("/health", handlers.HealthHandler)
	s.Router.Get("/readiness", handlers.ReadinessHandler)

	s.Router.Route("/api/v1", func(r chi.Router) {
		// Example of a route
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
}
