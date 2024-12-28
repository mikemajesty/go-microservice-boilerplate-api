package cache

import (
	"go-microservice-boilerplate-api/utils"
	"time"

	goCache "github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

type CacheAdapterType interface {
	~*redis.Client | ~*goCache.Cache
}

type CacheAdapter[T CacheAdapterType] interface {
	Connect() (T, utils.ApiException)
	Cache() T
	Delete(key string) utils.ApiException
	Get(key string) (string, utils.ApiException)
	Set(key string, value any, expired time.Duration) utils.ApiException
}
