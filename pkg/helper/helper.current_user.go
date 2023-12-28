package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CurrentUser(ctx *gin.Context) (string, error) {
	userId, exist := ctx.Get("userId")
	if exist == false {
		return "", fmt.Errorf("no user logged")
	}
	return fmt.Sprint(userId), nil
}
