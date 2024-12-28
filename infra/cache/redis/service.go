package cache

import (
	"context"
	"errors"
	"fmt"
	"go-microservice-boilerplate-api/infra/cache"
	"go-microservice-boilerplate-api/infra/secret"
	"time"

	"github.com/redis/go-redis/v9"
)

var EnvService = secret.SecretAdapter(secret.CreateSecret())

var ctx = context.Background()

var _cache *redis.Client

type adapter struct{}

func (a adapter) Delete(key string) error {
	err := _cache.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (a adapter) Set(key string, value any, expired time.Duration) error {
	err := _cache.Set(ctx, key, value, expired).Err()
	if err != nil {
		return err
	}

	return nil
}

func (a adapter) Get(key string) (string, error) {
	val, err := _cache.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
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

func CreateRedis() cache.CacheAdapter[*redis.Client] {
	return adapter{}
}
