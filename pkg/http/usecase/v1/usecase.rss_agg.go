package v1

import (
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
