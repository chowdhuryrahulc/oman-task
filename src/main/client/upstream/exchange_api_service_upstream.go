package upstream

import (
	"encoding/json"
	"fmt"

	"github.com/chowdhuryrahulc/oman-task/src/main/config"
	"github.com/chowdhuryrahulc/oman-task/src/main/models"
)

type IExchangeApiClient interface {
	GetExchangeApiResponse() (*models.ExchangeAPIResponse, error)
}

type ExchangeApiClient struct {
	eprModel models.IExchangeApiResponseModel
}

func NewExchangeApiClient(clientConfig *config.ExchangeApiConfig) *ExchangeApiClient {
	return &ExchangeApiClient{}
}

func (epsc *ExchangeApiClient) GetExchangeApiResponse() *models.ExchangeAPIResponse {
	byte := epsc.eprModel.Get()
	model := MarshalExchangeApiResponseToStruct(byte)
	return model
}

func MarshalExchangeApiResponseToStruct(res []byte) *models.ExchangeAPIResponse {
	exchangeApiResponse := &models.ExchangeAPIResponse{}

	err := json.Unmarshal(res, exchangeApiResponse)
	if err != nil {
		fmt.Println("error unmarshaling exchange api url response: " + err.Error())
	}
	fmt.Println("Exp", exchangeApiResponse, "StartDate", exchangeApiResponse.Rates.Dates.EUR)
	return exchangeApiResponse
}
