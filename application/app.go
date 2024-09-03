package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CristianSch248/CasaDoEstudanteDigital/routes"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: routes.LoadRutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Falha ao iniciar o serviço: %w", err)
	}

	return nil
}
