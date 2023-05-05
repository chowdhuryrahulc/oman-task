package config

import (
	"github.com/Netflix/go-env"
	"sync"
)

type Config struct {
	Port                          string `env:"PORT,default=4000"`
	ExchangeApiConfig *ExchangeApiConfig
	PostgressConfig              *PostgressConfig
}

var (
	appConfig     Config
	onceAppConfig sync.Once
)

func NewAppConfig() *Config {
	onceAppConfig.Do(func() {
		_, _ = env.UnmarshalFromEnviron(&appConfig)
		appConfig.PostgressConfig = NewPostgressConfig()
		appConfig.ExchangeApiConfig = NewExchangeApiConfig()
	})

	return &appConfig
}
