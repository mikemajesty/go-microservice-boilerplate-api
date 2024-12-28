package cache

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

type CacheAdapterType interface {
	~*redis.Client | ~*goCache.Cache
}

type CacheAdapter[T CacheAdapterType] interface {
	Connect() (T, error)
	Cache() T
	Delete(key string) error
	Get(key string) (string, error)
	Set(key string, value any, expired time.Duration) error
}
