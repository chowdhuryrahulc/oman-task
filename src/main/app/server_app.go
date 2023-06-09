package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/chowdhuryrahulc/oman-task/src/main/client/upstream"
	"github.com/chowdhuryrahulc/oman-task/src/main/config"
	"github.com/chowdhuryrahulc/oman-task/src/main/handler"
	"github.com/chowdhuryrahulc/oman-task/src/main/routes"
	"github.com/chowdhuryrahulc/oman-task/src/main/service"
)

func NewServer(appConfig *config.Config) (*http.Server, error) {
	exchangeApiClient := upstream.NewExchangeApiClient(appConfig.ExchangeApiConfig)
	exchangeApiService := service.NewExchangeApiService(exchangeApiClient)
	exchangeApiHandler := handler.NewExchangeApiHandler(exchangeApiService)

	return &http.Server{Addr: ":" + appConfig.Port, Handler: getRouter(
		// securityEnsurer,
		exchangeApiHandler)}, nil
}

func getRouter(
	// securityEnsurer *appMiddleware.SecurityEnsurer,
	exchangeApiHandler *handler.ExchangeApiHandler) *chi.Mux {

	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Compress(5),
		middleware.Recoverer,
	)

	routes.Configure(router,
		// securityEnsurer,
		exchangeApiHandler)

	return router
}
