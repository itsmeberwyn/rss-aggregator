package database

import (
	"context"
	"log"

	"github.com/itsmeberwyn/rss-service/config"
	"github.com/itsmeberwyn/rss-service/pkg/driver"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
for spawning multiple instance of connection
*/
func NewPostgresConnection() (*pgxpool.Pool, error) {
	connString := config.AppConfig.ConnString
	if config.AppConfig.Environment == "production" {
		connString = config.AppConfig.DsnConnString
	}
	conn, err := driver.InitializedPostgresConnection(connString)
	if err != nil {
		log.Fatal("Error creating connection ", err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return conn, nil
}
