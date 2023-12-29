package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
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

func UpdateLastFetchedFeed(conn *pgxpool.Pool, ctx context.Context, feed_id string) error {
	_, err := conn.Query(ctx,
		`UPDATE feeds SET
    last_fetched_at=$1
    WHERE id=$2
    `, time.Now(), feed_id)
	if err != nil {
		return err
	}
	return nil
}

func CreatePost(conn *pgxpool.Pool, ctx context.Context, post V1Model.PostModel) (V1Model.PostModel, error) {
	var postObj V1Model.PostModel
	err := conn.QueryRow(ctx,
		`INSERT INTO posts
    (id, created_at, updated_at, title, description, published_at, url, feed_id)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *
    `, uuid.New(), time.Now(), time.Now(), post.Title, post.Description, post.Published_at, post.Url, post.Feed_id).
		Scan(&postObj.Id, &postObj.Created_at, &postObj.Updated_at, &postObj.Title, &postObj.Description, &postObj.Published_at, &postObj.Url, &postObj.Feed_id)
	if err != nil {
		return postObj, err
	}
	return postObj, nil
}
