package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	handler "github.com/itsmeberwyn/rss-service/pkg/http/handler/v1"
	repository "github.com/itsmeberwyn/rss-service/pkg/http/repository/v1"
	usecase "github.com/itsmeberwyn/rss-service/pkg/http/usecase/v1"

	middleware "github.com/itsmeberwyn/rss-service/pkg/http/middleware"
)

type rssAggRoutes struct {
	db        *pgxpool.Pool
	router    *gin.RouterGroup
	handler   handler.RSSAggHandler
	middlware middleware.RSSAggMiddleware
}

func NewRSSAggRoutes(router *gin.RouterGroup, conn *pgxpool.Pool) *rssAggRoutes {
	// handler -> usecase -> repository architecture
	// building from low level to high level
	V1Repository := repository.NewRSSAggRepository(conn)
	V1UseCase := usecase.NewRSSAggUseCase(V1Repository)
	V1Handler := handler.NewRSSAggHandler(V1UseCase)
	V1Middleware := middleware.NewRSSAggMiddleware(V1UseCase)
	return &rssAggRoutes{
		db:        conn,
		router:    router,
		handler:   V1Handler,
		middlware: V1Middleware,
	}
}

func (r *rssAggRoutes) Routes() {
	V1Route := r.router.Group("/v1")
	{
		V1Route.POST("/user", r.handler.CreateUser)
		V1Route.GET("/ping", r.middlware.MiddlewareAuth(), func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": true,
				"method": "ping",
			})
		})
	}
}
