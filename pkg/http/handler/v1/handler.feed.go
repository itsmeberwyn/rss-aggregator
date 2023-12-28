package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itsmeberwyn/rss-service/pkg/helper"
	"github.com/itsmeberwyn/rss-service/pkg/http/handler"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
)

func (u RSSAggHandler) CreateFeed(ctx *gin.Context) {
	var feed *V1Model.FeedModel
	ctx.BindJSON(&feed)

	currentUser, err := helper.CurrentUser(ctx)
	if err != nil {
		handler.ErrorResponse(ctx, 401, err.Error())
		return
	}

	userId, err := uuid.Parse(currentUser)
	if err != nil {
		handler.ErrorResponse(ctx, 401, err.Error())
		return
	}

	feed.UserId = userId
	obj, statusCode, err := u.usecase.CreateFeed(ctx, feed)
	if err != nil {
		handler.ErrorResponse(ctx, statusCode, err.Error())
		return
	}
	handler.SuccessResponse(ctx, statusCode, "Success creating feed", obj)
}

func (u RSSAggHandler) GetUserFeeds(ctx *gin.Context) {
	currentUser, err := helper.CurrentUser(ctx)
	if err != nil {
		handler.ErrorResponse(ctx, 401, err.Error())
		return
	}

	userId, err := uuid.Parse(currentUser)
	if err != nil {
		handler.ErrorResponse(ctx, 401, err.Error())
		return
	}

	obj, statusCode, err := u.usecase.GetUserFeeds(ctx, fmt.Sprint(userId))
	if err != nil {
		handler.ErrorResponse(ctx, statusCode, err.Error())
		return
	}
	handler.SuccessResponse(ctx, statusCode, "Success getting user feeds", obj)
}
