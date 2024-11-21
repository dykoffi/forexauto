package config

import (
	"sync"

	"github.com/spf13/viper"
)

type ConfigInterface interface {
	New() *ConfigService
	Get(key string) string
	GetorThrow(key string) string
	GetOrDefault(key string, defaultValue string) string
}

type ConfigService struct {
	config *viper.Viper
}

// The unique instance for ConfigService (Singleton pattern)
var (
	iConfigService ConfigService
	once           sync.Once
)

func New() *ConfigService {

	once.Do(func() {
		iConfigService = ConfigService{
			config: viper.New(),
		}

		iConfigService.config.AutomaticEnv()
		iConfigService.config.SetConfigFile(".env")
		iConfigService.config.ReadInConfig()
	})

	return &iConfigService

}

func (cs *ConfigService) Get(key string) string {
	if value := cs.config.Get(key); value != nil {
		return value.(string)
	} else {
		return ""
	}
}

func (cs *ConfigService) GetOrThrow(key string) string {
	if value := cs.config.Get(key); value != nil {
		return value.(string)
	} else {
		panic(key + " doesn't exist in configuration env")
	}
}

func (cs *ConfigService) GetOrDefault(key string, defaultValue string) string {
	if value := cs.config.Get(key); value != nil {
		return value.(string)
	} else {
		return defaultValue
	}
}
