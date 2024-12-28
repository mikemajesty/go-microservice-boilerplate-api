package cache

import (
	"context"
	"fmt"
	"go-microservice-boilerplate-api/infra/cache"
	"go-microservice-boilerplate-api/infra/secret"
	"go-microservice-boilerplate-api/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

var EnvService = secret.SecretAdapter(secret.CreateSecret())

var ctx = context.Background()

var redisCache *redis.Client

type adapter struct{}

func (a adapter) Delete(key string) utils.ApiException {
	err := redisCache.Del(ctx, key).Err()
	if err != nil {
		return utils.ApiInternalServerException(err.Error())
	}

	return nil
}

func (a adapter) Set(key string, value any, expired time.Duration) utils.ApiException {
	err := redisCache.Set(ctx, key, value, expired).Err()
	if err != nil {
		return utils.ApiInternalServerException(err.Error())
	}

	return nil
}

func (a adapter) Get(key string) (string, utils.ApiException) {
	val, err := redisCache.Get(ctx, key).Result()
	if err != nil {
		return "", utils.ApiInternalServerException(err.Error())
	}

	return val, nil
}

func (a adapter) Cache() *redis.Client {
	return redisCache
}

func (a adapter) Connect() (*redis.Client, utils.ApiException) {
	host := EnvService.GetSecret("REDIS_HOST")
	port := EnvService.GetSecret("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, utils.ApiInternalServerException(err.Error())
	}

	redisCache = client

	fmt.Println("Successfully connected to Cache Redis")
	return client, nil
}

func CreateRedis() cache.CacheAdapter[*redis.Client] {
	return adapter{}
}
