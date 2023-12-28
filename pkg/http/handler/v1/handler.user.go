package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmeberwyn/rss-service/pkg/helper"
	"github.com/itsmeberwyn/rss-service/pkg/http/handler"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
)

func (u RSSAggHandler) CreateUser(ctx *gin.Context) {
	var user *V1Model.UserModel
	ctx.BindJSON(&user)
	obj, statusCode, err := u.usecase.CreateUser(ctx, user)
	if err != nil {
		handler.ErrorResponse(ctx, statusCode, err.Error())
		return
	}
	handler.SuccessResponse(ctx, statusCode, "Success creating user", obj)
}

func (u RSSAggHandler) GetUserByAPIKey(ctx *gin.Context) {
	apiKey, err := helper.ExtractAPIKey(ctx.Request.Header["Authorization"][0])
	if err != nil {
		handler.ErrorResponse(ctx, 400, err.Error())
		return
	}
	obj, statusCode, err := u.usecase.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		handler.ErrorResponse(ctx, statusCode, err.Error())
		return
	}
	handler.SuccessResponse(ctx, statusCode, "Success getting user data", obj)
}
