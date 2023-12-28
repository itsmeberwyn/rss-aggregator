package v1

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type RSSAggRepository struct {
	conn *pgxpool.Pool
}

func NewRSSAggRepository(conn *pgxpool.Pool) RSSAggRepository {
	return RSSAggRepository{
		conn: conn,
	}
}
