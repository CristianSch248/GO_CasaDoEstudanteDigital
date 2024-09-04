package routes

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func LoadRutes() *chi.Mux {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/usuarios", LoadUserRoutes)

	return router
}
