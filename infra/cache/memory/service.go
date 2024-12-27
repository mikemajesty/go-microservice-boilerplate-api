package cache

import (
	cacheAdapter "go-microservice-boilerplate-api/infra/cache"

	"github.com/patrickmn/go-cache"
)

type adapter struct{}

var _cache *cache.Cache

func (a adapter) Cache() *cache.Cache {
	return _cache
}

func (a adapter) Connect() (*cache.Cache, error) {
	return _cache, nil
}

func CreateRedis() cacheAdapter.CacheAdapter[*cache.Cache] {
	return adapter{}
}
