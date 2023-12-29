package server

/**
  Create server that can spawn multiple instance
**/

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsmeberwyn/rss-service/config"
	"github.com/itsmeberwyn/rss-service/pkg/database"
	"github.com/itsmeberwyn/rss-service/pkg/helper"
	"github.com/itsmeberwyn/rss-service/pkg/routes"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	conn, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	go helper.StartScraping(conn, 10, time.Minute)

	router := NewRouter()
	api := router.Group("api")
	api.GET("/", routes.RootRoutes)
	routes.NewRSSAggRoutes(api, conn).Routes()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler: router,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() error {
	log.Printf("Server starting on port%v", a.HttpServer.Addr)
	err := a.HttpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("failed to listen and server")
		return err
	}
	return nil
}

func NewRouter() *gin.Engine {
	mode := gin.ReleaseMode
	if config.AppConfig.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	router := gin.Default()
	return router
}
