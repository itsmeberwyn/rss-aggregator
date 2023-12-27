package handler

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, BaseResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, BaseResponse{
		Status:  false,
		Message: message,
	})
}

func AbortResponse(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, BaseResponse{
		Status:  false,
		Message: message,
	})
}

func RedirectResponse(ctx *gin.Context, location string) {
	redirect := url.URL{Path: location}
	ctx.Redirect(http.StatusFound, redirect.RequestURI())
}
