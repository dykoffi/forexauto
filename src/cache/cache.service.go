package cache

import (
	"sync"
	"time"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/go-redis/redis"
)

type CacheInterface interface {
	Set(key string, value any)
	Get(key string) (any, bool)
}

type CacheService struct {
	cs  *redis.Client
	ttl time.Duration
}

// The unique instance for CacheService (Singleton pattern)
var (
	iCacheService CacheService
	once          sync.Once
)

func New(config config.ConfigInterface) *CacheService {

	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr: config.GetOrThrow("REDIS_ADDR"),
		})

		err := client.Ping().Err()
		if err != nil {
			panic(err)
		}
		iCacheService = CacheService{cs: client}
	})

	return &iCacheService

}

func (c *CacheService) Set(key string, value any) {
	c.cs.Set(key, value, c.ttl)
}

func (c *CacheService) Get(key string) (any, bool) {
	s, err := c.cs.Get(key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
