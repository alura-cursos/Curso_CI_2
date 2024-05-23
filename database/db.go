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

// ConectaComBancoDeDados initializes the database connection
func ConectaComBancoDeDados() {
	stringDeConexao := "host=" + os.Getenv("HOST") + 
					   " user=" + os.Getenv("USER") + 
					   " password=" + os.Getenv("PASSWORD") + 
					   " dbname=" + os.Getenv("DBNAME") + 
					   " port=" + os.Getenv("PORT") + 
					   " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Panic("Erro ao realizar a migração: ", err)
	}
}

// FechaConexaoComBancoDeDados closes the database connection
func FechaConexaoComBancoDeDados() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Panic("Erro ao obter a instância de sql.DB: ", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Panic("Erro ao fechar a conexão com banco de dados: ", err)
	}
}
