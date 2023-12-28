package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmeberwyn/rss-service/pkg/helper"
	"github.com/itsmeberwyn/rss-service/pkg/http/handler"
	usecase "github.com/itsmeberwyn/rss-service/pkg/http/usecase/v1"
)

type RSSAggMiddleware struct {
	usecase usecase.RSSAggUseCase
}

func NewRSSAggMiddleware(usecase usecase.RSSAggUseCase) RSSAggMiddleware {
	return RSSAggMiddleware{
		usecase: usecase,
	}
}

func (u RSSAggMiddleware) MiddlewareAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		AuthHeader := ctx.Request.Header["Authorization"]
		if len(AuthHeader) == 0 {
			handler.AbortResponse(ctx, "no headers provided")
			return
		}

		apiKey, err := helper.ExtractAPIKey(AuthHeader[0])
		if err != nil {
			handler.AbortResponse(ctx, err.Error())
			return
		}

		_, statusCode, err := u.usecase.GetUserByAPIKey(ctx, apiKey)
		if err != nil {
			handler.ErrorResponse(ctx, statusCode, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
