package server

/**
  Create server that can spawn multiple instance
**/

import (
	"log"
	"net/http"

	"github.com/itsmeberwyn/rss-service/config"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	server := &http.Server{
		Addr: ":" + config.AppConfig.Port,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() error {
	err := a.HttpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("failed to listen and server")
	}
	return nil
}
