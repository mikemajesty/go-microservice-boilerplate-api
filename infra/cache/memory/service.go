package cache

import (
	"fmt"
	cacheAdapter "go-microservice-boilerplate-api/infra/cache"

	"github.com/patrickmn/go-cache"
)

type adapter struct{}

var _cache *cache.Cache

func (a adapter) Cache() *cache.Cache {
	return _cache
}

func (a adapter) Connect() (*cache.Cache, error) {
	fmt.Println("Successfully connected to Cache Memory")
	return _cache, nil
}

func CreateMemory() cacheAdapter.CacheAdapter[*cache.Cache] {
	return adapter{}
}
