package v1

import (
	"github.com/go-chi/chi"
	"github.com/chowdhuryrahulc/oman-task/src/main/handler"
)

// IncludeExchangeApiRoute ...
func IncludeExchangeApiRoute(
	router chi.Router,
	exchangeApiHandler *handler.ExchangeApiHandler) chi.Router {

	router.Get("/", exchangeApiHandler.GetExchangeApiResponse)
	return router
}