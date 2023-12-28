package v1

import (
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
