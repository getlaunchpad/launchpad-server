package user

import (
	"net/http"

	"github.com/lucasstettner/launchpad-server/app/utils/responses"

	"github.com/lucasstettner/launchpad-server/app/models"
	"github.com/lucasstettner/launchpad-server/app/utils/jwt"

	"github.com/go-chi/chi"
	"github.com/lucasstettner/launchpad-server/config"
)

type Config struct {
	*config.Config
}

func New(c *config.Config) *Config {
	return &Config{c}
}

func (c *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/me", c.meRoute)
	return router
}

func (c *Config) meRoute(w http.ResponseWriter, r *http.Request) {
	payload := jwt.Token{}.New().Decode(r)

	u := &models.User{}
	if err := u.FindUserByID(c.DB, payload.UserID); err != nil {
		responses.NewResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	responses.NewResponse(w, http.StatusOK, nil, u)
}
