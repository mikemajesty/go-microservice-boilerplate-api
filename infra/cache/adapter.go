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
	Connect() (T, *utils.AppException)
	Cache() T
	Delete(key string) *utils.AppException
	Get(key string) (string, *utils.AppException)
	Set(key string, value any, expired time.Duration) *utils.AppException
}
