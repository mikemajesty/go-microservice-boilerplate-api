package main

import (
	"go-microservice-boilerplate-api/infra/cache"
	redisCache "go-microservice-boilerplate-api/infra/cache/redis"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	"go-microservice-boilerplate-api/infra/secret"
	modules_cat "go-microservice-boilerplate-api/modules/cat"
	modules_dog "go-microservice-boilerplate-api/modules/dog"
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var SecretService = secret.Adapter(secret.CreateSecret())
var PostgresService = infra_database.Adapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())
var MongoService = infra_database.Adapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())
var RedisService = cache.CacheAdapter[*redis.Client](redisCache.CreateRedis())

func init() {
	SecretService.InitEnvs()
	PostgresService.Connect()
	MongoService.Connect()
	RedisService.Connect()
}

func main() {
	utils.Route.GET("/", func(c *gin.Context) {
		c.JSON(200, "up and running")
	})
	modules_cat.CreateCatRoutes()
	modules_dog.CreateDogRoutes()

	utils.Route.Run()
}
