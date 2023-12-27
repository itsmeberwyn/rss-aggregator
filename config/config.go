package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/*
Singleton method where we can share this variable
throughout the whole project
*/
var AppConfig appConfig

// structure of the config data from .env
type appConfig struct {
	Port        int
	Environment string
	Debug       bool
	ConnString  string
}

func InitializeAppConfig() error {
	godotenv.Load(".env")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return fmt.Errorf("error parsing port value")
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		return fmt.Errorf("environment is empty")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return fmt.Errorf("error parsing debug value")
	}

	connString := os.Getenv("CONNSTRING")
	if connString == "" {
		return fmt.Errorf("connection string is empty")
	}

	AppConfig = appConfig{
		Port:        port,
		Environment: environment,
		Debug:       debug,
		ConnString:  connString,
	}
	return nil
}
