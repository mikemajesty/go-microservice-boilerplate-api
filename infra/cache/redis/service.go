package cache

import (
	"context"
	"errors"
	"fmt"
	"go-microservice-boilerplate-api/infra/cache"
	"go-microservice-boilerplate-api/infra/secret"

	"github.com/redis/go-redis/v9"
)

var EnvService = secret.Adapter(secret.CreateSecret())

var ctx = context.Background()

var _cache *redis.Client

type adapter struct{}

func CreateRedis() cache.CacheAdapter[*redis.Client] {
	return adapter{}
}

func (a adapter) Cache() *redis.Client {
	return _cache
}

func (a adapter) Connect() (*redis.Client, error) {
	host := EnvService.GetSecret("REDIS_HOST")
	port := EnvService.GetSecret("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, errors.New("failed to connect to Redis")
	}

	_cache = client

	fmt.Println("Successfully connected to Cache Redis")
	return client, nil
}
