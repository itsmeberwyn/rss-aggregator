package v1

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
)

func (r *RSSAggRepository) CreateFeed(ctx *gin.Context, feed *V1Model.FeedModel) (V1Model.FeedModel, error) {
	var feedObj V1Model.FeedModel
	err := r.conn.QueryRow(ctx,
		`INSERT INTO feeds 
    (id, created_at, updated_at, name, url, user_id) 
    VALUES ($1, $2, $3, $4, $5, $6) 
    RETURNING *
    `, uuid.New(), time.Now(), time.Now(), feed.Name, feed.Url, feed.UserId).
		Scan(&feedObj.Id, &feedObj.Created_at, &feedObj.Updated_at, &feedObj.Name, &feedObj.Url, &feedObj.UserId, &feedObj.Last_fetched_at)
	if err != nil {
		return feedObj, err
	}
	return feedObj, nil
}

func (r *RSSAggRepository) GetUserFeeds(ctx *gin.Context, user_id string) ([]V1Model.FeedModel, error) {
	var feeds []V1Model.FeedModel
	rows, err := r.conn.Query(ctx,
		`SELECT * FROM feeds
    WHERE user_id=$1
    `, user_id)
	if err != nil {
		return feeds, err
	}
	defer rows.Close()

	for rows.Next() {
		var feed V1Model.FeedModel
		err := rows.Scan(&feed.Id, &feed.Created_at, &feed.Updated_at, &feed.Name, &feed.Url, &feed.UserId, &feed.Last_fetched_at)
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
