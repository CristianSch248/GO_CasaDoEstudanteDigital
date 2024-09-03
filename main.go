package main

import (
	"context"
	"fmt"

	"github.com/CristianSch248/CasaDoEstudanteDigital/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Falha ao iniciar o App: ", err)
	}
}
