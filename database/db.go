package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func envOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	if value, ok := os.LookupEnv(strings.ToLower(key)); ok && value != "" {
		return value
	}
	return fallback
}

func ConectaComBancoDeDados() {
	host := envOrDefault("HOST", "localhost")
	user := envOrDefault("USER", "root")
	password := envOrDefault("PASSWORD", "root")
	dbname := envOrDefault("DBNAME", "root")
	port := envOrDefault("PORT", "5432")

	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	DB.AutoMigrate(&models.Aluno{})
}
