package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/cors"

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
		corsConfig().Handler,       // Sets up cors for use in production
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responses.Success(w, http.StatusOK, "Hello Launchpad!")
	})

	// Mount routes on endpoint /VERSION_NUMBER/...
	router.Route(fmt.Sprintf("/%s", c.Constants.Version), func(r chi.Router) {
		r.Mount("/status", status.Routes())
		r.Mount("/auth/google", auth.New(c).Routes())
	})

	return router
}

func corsConfig() *cors.Cors {
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	return cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	})
}
