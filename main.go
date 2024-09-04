package main

import (
	"context"
	"fmt"

	"github.com/CristianSch248/CasaDoEstudanteDigital/application"
	"github.com/CristianSch248/CasaDoEstudanteDigital/configs"
)

func main() {

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	app := application.New()

	err = app.Start(context.TODO())
	if err != nil {
		fmt.Println("Falha ao iniciar o App: ", err)
	}
}
