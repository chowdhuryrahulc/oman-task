package routes

import (
	"github.com/go-chi/chi"
	"github.com/chowdhuryrahulc/oman-task/src/main/routes/v1"
	"github.com/chowdhuryrahulc/oman-task/src/main/handler"
)

func Configure(router *chi.Mux,
	// securityEnsurer *middleware.SecurityEnsurer,
	exchangeApiHandler *handler.ExchangeApiHandler) chi.Router {

	// Routes
	router.Group(func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			v1.IncludeExchangeApiRoute(r, exchangeApiHandler)
		})
	})
 
	return router
}
