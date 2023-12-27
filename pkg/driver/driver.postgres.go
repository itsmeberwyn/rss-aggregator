package driver

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializedPostgresConnection(connString string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("error initializing postgres connection %v", err)
	}

	return dbpool, nil
}
