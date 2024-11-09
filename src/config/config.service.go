package config

import (
	"github.com/spf13/viper"
)

type ConfigInterface interface {
	New() *ConfigService
	Get(key string) string
	GetorThrow(key string) string
	GetOrDefault(key string, defaultValue string)
}

type ConfigService struct {
	config *viper.Viper
}

// The unique instance for ConfigService (Singleton pattern)
var IConfigService ConfigService

func New() *ConfigService {

	if (IConfigService != ConfigService{}) {
		return &IConfigService
	}

	IConfigService = ConfigService{
		config: viper.New(),
	}

	IConfigService.config.AutomaticEnv()
	IConfigService.config.SetConfigFile(".env")
	IConfigService.config.ReadInConfig()

	return &IConfigService

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
