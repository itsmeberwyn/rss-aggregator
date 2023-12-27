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
	conn, err := driver.InitializedPostgresConnection(config.AppConfig.ConnString)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return conn, nil
}
