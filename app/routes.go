package app

import (
	"github.com/lucasstettner/launchpad-server/app/handlers"

	"github.com/go-chi/chi"
)

func (s *Server) initializeRoutes() {
	s.Router.Route("/api/v1", func(r chi.Router) {
		// These routes are used for kubernetes
		r.Get("/health", handlers.HealthHandler)
		r.Get("/readiness", handlers.ReadinessHandler)
	})
}
