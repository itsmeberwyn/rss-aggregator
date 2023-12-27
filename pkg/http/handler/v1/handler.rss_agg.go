package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmeberwyn/rss-service/pkg/http/handler"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
	usecase "github.com/itsmeberwyn/rss-service/pkg/http/usecase/v1"
)

type RSSAggHandler struct {
	usecase usecase.RSSAggUseCase
}

func NewRSSAggHandler(usecase usecase.RSSAggUseCase) RSSAggHandler {
	return RSSAggHandler{
		usecase: usecase,
	}
}

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
