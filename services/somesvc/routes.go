package usersvc

import (
	"github.com/go-chi/chi"
	_ "github.com/gpng/ring-fit-maker/docs" // required for generating docs
)

// Routes for app
func (s *Service) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", s.handleStatus())

	return router
}
