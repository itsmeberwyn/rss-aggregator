package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	ApiKey     string    `json:"apikey"`
}
