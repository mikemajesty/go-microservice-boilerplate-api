package cache

import (
	"fmt"
	cacheAdapter "go-microservice-boilerplate-api/infra/cache"
	"time"

	"github.com/patrickmn/go-cache"
)

type adapter struct{}

var _cache *cache.Cache

func (a adapter) Cache() *cache.Cache {
	return _cache
}

func (a adapter) Connect() (*cache.Cache, error) {
	fmt.Println("Successfully connected to Cache Memory")
	_cache = cache.New(cache.DefaultExpiration, cache.DefaultExpiration)
	return _cache, nil
}

func (a adapter) Set(key string, value any, expired time.Duration) error {
	_cache.Set(key, value, cache.DefaultExpiration)
	return nil
}

func (a adapter) Get(key string) (string, error) {
	val, found := _cache.Get(key)
	if !found {
		return "", nil
	}

	return val.(string), nil
}

func (a adapter) Delete(key string) error {
	_cache.Delete(key)
	return nil
}

func CreateMemory() cacheAdapter.CacheAdapter[*cache.Cache] {
	return adapter{}
}
