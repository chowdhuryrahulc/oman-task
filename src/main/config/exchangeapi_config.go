package config

import (
	"sync"

	"github.com/Netflix/go-env"
)

type ExchangeApiConfig struct {
	Host               string `env:"EXCHANGE_API_HOST,default=http://localhost:8199"`
	ExchangeApiTimeout int    `env:"EXCHANGE_API_TIMEOUT,default=500"`
	AuthToken          string `env:"EXCHANGE_API_AUTH_TOKEN,default=Bearer 39852a79-0c49-40f6-9dec-0247048d631d"`
}

var (
	exchangeApiConfig     ExchangeApiConfig
	onceExchangeApiConfig sync.Once
)

// NewExchangeApiConfig is singleton and it makes the configs exportable
func NewExchangeApiConfig() *ExchangeApiConfig {
	onceExchangeApiConfig.Do(func() {
		_, _ = env.UnmarshalFromEnviron(&exchangeApiConfig)
	})

	return &exchangeApiConfig
}
