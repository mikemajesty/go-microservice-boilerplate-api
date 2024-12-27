package cache

import (
	goCache "github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

type CacheAdapterType interface {
	~*redis.Client | ~*goCache.Cache
}

type CacheAdapter[T CacheAdapterType] interface {
	Connect() (T, error)
	Cache() T
}
