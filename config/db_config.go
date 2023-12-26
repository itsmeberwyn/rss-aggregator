package config

import (
	"os"

	"github.com/joho/godotenv"
)

var DBConfig DbConfig

type DbConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func InitializeDBConfig() error {
	godotenv.Load(".env")
	DBConfig = DbConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_DATABASE"),
	}
	return nil
}
