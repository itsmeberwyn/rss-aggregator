package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
	repository "github.com/itsmeberwyn/rss-service/pkg/http/repository/v1"
)

type RSSAggUseCase struct {
	repository repository.RSSAggRepository
}

func NewRSSAggUseCase(repository repository.RSSAggRepository) RSSAggUseCase {
	return RSSAggUseCase{
		repository: repository,
	}
}

func (u RSSAggUseCase) CreateUser(ctx *gin.Context, user *V1Model.UserModel) (V1Model.UserModel, int, error) {
	obj, err := u.repository.CreateUser(ctx, user)
	if err != nil {
		return obj, 500, fmt.Errorf("error creating new user %v", err)
	}
	return obj, 200, nil
}

func (u RSSAggUseCase) GetUserByAPIKey(ctx *gin.Context, apiKey string) (V1Model.UserModel, int, error) {
	obj, err := u.repository.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		return obj, 404, fmt.Errorf("no user found %v", err)
	}
	return obj, 200, nil
}
