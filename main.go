package main

import (
	"context"
	"go-microservice-boilerplate-api/infra/cache"
	memoryCache "go-microservice-boilerplate-api/infra/cache/memory"
	redisCache "go-microservice-boilerplate-api/infra/cache/redis"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	infra "go-microservice-boilerplate-api/infra/logger"
	"go-microservice-boilerplate-api/infra/secret"
	modules_cat "go-microservice-boilerplate-api/modules/cat"
	modules_dog "go-microservice-boilerplate-api/modules/dog"
	"go-microservice-boilerplate-api/observables"
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	cacheMemory "github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var SecretService = secret.SecretAdapter(secret.CreateSecret())
var PostgresService = infra_database.DatabaseAdapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())
var MongoService = infra_database.DatabaseAdapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())
var RedisService = cache.CacheAdapter[*redis.Client](redisCache.CreateRedis())
var CacheMemoryService = cache.CacheAdapter[*cacheMemory.Cache](memoryCache.CreateMemory())
var LoggerService = infra.LoggerAdapter(infra.CreateLogger())

func init() {
	SecretService.InitEnvs()
	PostgresService.Connect()
	MongoService.Connect()
	RedisService.Connect()
	CacheMemoryService.Connect()
	mw := &infra.MongoWriter{Client: MongoService.DB()}
	LoggerService.Connect(mw)
}

func main() {
	ctx := context.Background()

	utils.Route.Use(observables.GinBodyLogMiddleware(LoggerService))
	utils.Route.Use(func(g *gin.Context) {
		id := uuid.New().String()
		g.Set("traceId", id)
		LoggerService.SetContext(utils.ContextWithValues(ctx, "traceId", id))
		g.Next()
	})

	utils.Route.GET("/", func(c *gin.Context) {
		c.JSON(200, "up and running")
	})

	modules_cat.CreateCatRoutes()
	modules_dog.CreateDogRoutes()

	utils.Route.Run()
}
