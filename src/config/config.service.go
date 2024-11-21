package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Interface interface {
	Get(key string) string
	GetOrThrow(key string) string
	GetOrDefault(key string, defaultValue string) string
}

type Service struct {
	config *viper.Viper
}

// The unique instance for ConfigService (Singleton pattern)
var (
	iConfigService Service
	once           sync.Once
)

func New() *Service {

	once.Do(func() {
		iConfigService = Service{
			config: viper.New(),
		}

		iConfigService.config.AutomaticEnv()
		iConfigService.config.SetConfigFile(".env")
		iConfigService.config.ReadInConfig()
	})

	return &iConfigService

}

func (cs *Service) Get(key string) string {
	if value := cs.config.Get(key); value != nil {
		return value.(string)
	} else {
		return ""
	}
}

func (cs *Service) GetOrThrow(key string) string {
	if value := cs.config.Get(key); value != nil {
		return value.(string)
	} else {
		panic(key + " doesn't exist in configuration env")
	}
}

func (cs *Service) GetOrDefault(key string, defaultValue string) string {
	if value := cs.config.Get(key); value != nil {
		return value.(string)
	} else {
		return defaultValue
	}
}
