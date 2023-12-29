package repository

import (
	"context"

	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func FetchFeed(conn *pgxpool.Pool, ctx context.Context, limit int32) ([]V1Model.FeedModel, error) {
	var feeds []V1Model.FeedModel
	rows, err := conn.Query(ctx,
		`
    SELECT * FROM feeds
    ORDER BY last_fetched_at
    NULLS FIRST
  `)
	if err != nil {
		return feeds, err
	}

	for rows.Next() {
		var feed V1Model.FeedModel
		err = rows.Scan(&feed.Id, &feed.Created_at, &feed.Updated_at, &feed.Name, &feed.Url, &feed.UserId, &feed.Last_fetched_at)
		if err != nil {
			return feeds, err
		}
		feeds = append(feeds, feed)
	}

	if err := rows.Err(); err != nil {
		return feeds, err
	}
	return feeds, nil
}
