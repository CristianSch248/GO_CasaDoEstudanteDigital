package main

import (
	"log"

	"github.com/CristianSch248/CasaDoEstudanteDigital/configs"
	"github.com/CristianSch248/CasaDoEstudanteDigital/db"
	"github.com/CristianSch248/CasaDoEstudanteDigital/models"
)

func main() {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações: %v", err)
	}

	// Abrir a conexão com o banco de dados
	connection, err := db.OpenConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	// Executar a migração
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
