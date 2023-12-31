package v1

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
)

func (r *RSSAggRepository) CreateUser(ctx *gin.Context, user *V1Model.UserModel) (V1Model.UserModel, error) {
	var userObj V1Model.UserModel
	err := r.conn.QueryRow(ctx,
		`INSERT INTO users 
    (id, created_at, updated_at, name) 
    VALUES ($1, $2, $3, $4) 
    RETURNING *
    `, uuid.New(), time.Now(), time.Now(), user.Name).
		Scan(&userObj.Id, &userObj.Created_at, &userObj.Updated_at, &userObj.Name, &userObj.ApiKey)
	if err != nil {
		return userObj, err
	}
	return userObj, nil
}

func (r *RSSAggRepository) GetUserByAPIKey(ctx *gin.Context, apiKey string) (V1Model.UserModel, error) {
	var userObj V1Model.UserModel
	err := r.conn.QueryRow(ctx,
		`SELECT * FROM users 
    WHERE api_key=$1
    `, apiKey).
		Scan(&userObj.Id, &userObj.Created_at, &userObj.Updated_at, &userObj.Name, &userObj.ApiKey)
	if err != nil {
		return userObj, err
	}
	return userObj, nil
}
