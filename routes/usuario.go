package routes

import (
	"github.com/CristianSch248/CasaDoEstudanteDigital/handlers"
	"github.com/go-chi/chi/v5"
)

func LoadUserRoutes(router chi.Router) {
	userHandler := &handlers.User{}

	router.Post("/new", userHandler.Create)
	router.Get("/get", userHandler.Read)
	router.Put("/update", userHandler.Update)
	router.Delete("/delete", userHandler.Delete)
}
