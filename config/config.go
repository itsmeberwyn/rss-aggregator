package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig appConfig

// structure of the config data from .env
type appConfig struct {
	Port        string
	Environment string
	Debug       string
}

func InitializeAppConfig() error {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	environment := os.Getenv("ENVIRONMENT")
	debug := os.Getenv("DEBUG")

	if port == "" || environment == "" {
		return fmt.Errorf("Config port or environment is empty!")
	}

	AppConfig = appConfig{
		Port:        port,
		Environment: environment,
		Debug:       debug,
	}
	return nil
}
