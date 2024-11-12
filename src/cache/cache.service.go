package cache

import (
	"context"
	"time"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/go-redis/redis"
)

type CacheService struct {
	cs  *redis.Client
	ttl time.Duration
}

// The unique instance for CacheService (Singleton pattern)
var ICacheService CacheService

func New() *CacheService {

	if (ICacheService != CacheService{}) {
		return &ICacheService
	}

	config := config.New()

	client := redis.NewClient(&redis.Options{
		Addr: config.GetOrThrow("REDIS_ADDR"),
	})

	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}

	ICacheService = CacheService{cs: client}

	return &ICacheService

}

func (c *CacheService) Set(ctx context.Context, key string, value interface{}) {
	c.cs.Set(key, value, c.ttl)
}

func (c *CacheService) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.cs.Get(key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
