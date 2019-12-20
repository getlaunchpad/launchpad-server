package status

import (
	"fmt"
	"net/http"

	"github.com/lucasstettner/launchpad-server/app/utils/jwt"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/health", healthHandler)
	router.Get("/readiness", readinessHandler)
	return router
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	token := &jwt.Token{}
	fmt.Println(token.New().Encode(1, "member"))
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
