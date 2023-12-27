package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmeberwyn/rss-service/pkg/http/handler"
)

func RootRoutes(ctx *gin.Context) {
	handler.SuccessResponse(ctx, 200, "RSS Aggregator", struct{}{})
}
