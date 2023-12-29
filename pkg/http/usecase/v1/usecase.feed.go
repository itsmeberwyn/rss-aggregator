package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
)

func (u RSSAggUseCase) CreateFeed(ctx *gin.Context, feed *V1Model.FeedModel) (V1Model.FeedModel, int, error) {
	obj, err := u.repository.CreateFeed(ctx, feed)
	if err != nil {
		return obj, 400, fmt.Errorf("error creating new feed %v", err)
	}
	return obj, 200, nil
}

func (u RSSAggUseCase) GetUserFeeds(ctx *gin.Context, user_id string) ([]V1Model.FeedModel, int, error) {
	obj, err := u.repository.GetUserFeeds(ctx, user_id)
	if err != nil {
		return obj, 400, fmt.Errorf("error getting user feeds %v", err)
	}
	return obj, 200, nil
}

func (u RSSAggUseCase) GetFeedById(ctx *gin.Context, feed_id string) (V1Model.FeedModel, int, error) {
	obj, err := u.repository.GetFeedById(ctx, feed_id)
	if err != nil {
		return obj, 400, fmt.Errorf("no feed found %v", err)
	}
	return obj, 200, nil
}

func (u RSSAggUseCase) DeleteUserFeed(ctx *gin.Context, user_id string, feed_id string) error {
	err := u.repository.DeleteUserFeed(ctx, user_id, feed_id)
	if err != nil {
		return err
	}
	return nil
}
