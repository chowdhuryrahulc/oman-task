package handler

import (
	"net/http"

	"github.com/chowdhuryrahulc/oman-task/src/main/service"
	"github.com/chowdhuryrahulc/oman-task/src/main/util"
)

// ExchangeApiHandler ...
type ExchangeApiHandler struct {
	exchangeApiService service.IExchangeApiService
}

// NewExchangeApiHandler ...
func NewExchangeApiHandler(
	exchangeApiService service.IExchangeApiService) *ExchangeApiHandler {

	return &ExchangeApiHandler{
		exchangeApiService: exchangeApiService,
	}
}

// GetExchangeApiResponse ...
func (c *ExchangeApiHandler) GetExchangeApiResponse(w http.ResponseWriter, r *http.Request) {
	res, err := c.exchangeApiService.GetExchangeApiResponse()
	if err != nil {
		util.SendInternalError(w, r, "Internal Error", err)
		return
	}

	util.Send(w, r, "Success", res)
}
