package main

import (
	"log"

	"github.com/CristianSch248/CasaDoEstudanteDigital/db"
	"github.com/CristianSch248/CasaDoEstudanteDigital/models"
)

func main() {
	connection, err := db.OpenConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	err = connection.AutoMigrate(
		&models.Usuario{},
		&models.Vaga{},
		&models.Apartamento{},
		&models.Manutencao{},
		&models.Patrimonio{},
		&models.Vistoria{},
	)
	if err != nil {
		log.Fatalf("Erro ao migrar o banco de dados: %v", err)
	}

	log.Println("Migração concluída com sucesso")
}
