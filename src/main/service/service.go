package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/chowdhuryrahulc/oman-task/src/main/models"
	"github.com/chowdhuryrahulc/oman-task/src/main/client/upstream"
)

// IExchangeApiService ...
type IExchangeApiService interface {
	GetExchangeApiResponse() (*models.ExchangeAPIResponse, error)
}

type ExchangeApiService struct {
	epsClient upstream.IExchangeApiServiceClient
}

func NewExchangeApiService(epsClient upstream.IExchangeApiServiceClient) *ExchangeApiService {
	return &ExchangeApiService{
		epsClient: epsClient,
	}
}

func (cs *ExchangeApiService) GetExchangeApiResponse() (*models.ExchangeAPIResponse, error) {
	res, err := cs.epsClient.GetExchangeApiResponse()
	if err != nil {
		log.Errorf("ExchangeApiService: SendExchangeApiResponse API error: %s err: %v", err)
		return nil, err
	}

	log.Infof("ExchangeApiService: SendExchangeApiResponse API: %s, response: %v", res)
	return res, nil
}
