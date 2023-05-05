package config

import (
	"github.com/Netflix/go-env"
	"sync"
)

// PostgressConfig ...
type PostgressConfig struct {
	// 
}

var (
	postgressConfig     PostgressConfig
	onceMiddlewareConfig sync.Once
)

// NewPostgressConfig is singleton and it makes the configs exportable
func NewPostgressConfig() *PostgressConfig {
	onceMiddlewareConfig.Do(func() {
		_, _ = env.UnmarshalFromEnviron(&postgressConfig)
	})

	return &postgressConfig
}
