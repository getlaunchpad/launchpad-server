package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/lucasstettner/launchpad-server/app/features/auth"
	"github.com/lucasstettner/launchpad-server/app/features/status"
	"github.com/lucasstettner/launchpad-server/app/utils/responses"
	"github.com/lucasstettner/launchpad-server/config"
)

func Routes(c *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responses.Success(w, http.StatusOK, "Hello Launchpad!")
	})

	// Mount routes on endpoint /v1/...
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/status", status.Routes())
		r.Mount("/auth/google", auth.New(c).Routes())
	})

	return router
}
