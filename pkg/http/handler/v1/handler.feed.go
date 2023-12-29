package v1

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itsmeberwyn/rss-service/pkg/helper"
	"github.com/itsmeberwyn/rss-service/pkg/http/handler"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
)

func (h RSSAggHandler) CreateFeed(ctx *gin.Context) {
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

	url, err := url.Parse(feed.Url)
	fmt.Println(url.Hostname())
	if url.Hostname() != "www.youtube.com" {
		handler.ErrorResponse(ctx, 401, fmt.Sprint("invalid url (youtube links only)"))
		return
	}
	if err != nil {
		handler.ErrorResponse(ctx, 401, err.Error())
		return
	}

	feed.UserId = userId
	obj, statusCode, err := h.usecase.CreateFeed(ctx, feed)
	if err != nil {
		handler.ErrorResponse(ctx, statusCode, err.Error())
		return
	}
	handler.SuccessResponse(ctx, statusCode, "Success creating feed", obj)
}

func (h RSSAggHandler) GetUserFeeds(ctx *gin.Context) {
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

	obj, statusCode, err := h.usecase.GetUserFeeds(ctx, fmt.Sprint(userId))
	if err != nil {
		handler.ErrorResponse(ctx, statusCode, err.Error())
		return
	}
	handler.SuccessResponse(ctx, statusCode, "Success getting user feeds", obj)
}

func (h RSSAggHandler) DeleteUserFeed(ctx *gin.Context) {
	feedId := ctx.Param("feedId")
	obj, statusCode, err := h.usecase.GetFeedById(ctx, feedId)
	if err != nil {
		handler.ErrorResponse(ctx, statusCode, err.Error())
		return
	}

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

	err = h.usecase.DeleteUserFeed(ctx, fmt.Sprint(userId), fmt.Sprint(obj.Id))
	if err != nil {
		handler.ErrorResponse(ctx, 400, err.Error())
	}

	handler.SuccessResponse(ctx, statusCode, "Success deleting feed", struct{}{})
}
