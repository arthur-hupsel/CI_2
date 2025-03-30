package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	os.Getenv("HOST")

	stringDeConexao := "host="+OS.Getenv("HOST")+" user="+OS.Getenv("USER")+" password="+OS.Getenv("PASSWORD")+" dbname="+OS.Getenv("DBNAME")+" port="+OS.Getenv("PORT")+" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
