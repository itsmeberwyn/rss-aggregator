package main

import (
	"log"

	server "github.com/itsmeberwyn/rss-service/cmd/api"
	"github.com/itsmeberwyn/rss-service/config"
)

// initialize configs
func init() {
	err := config.InitializeAppConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app, err := server.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
