package main

import (
	"fmt"
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	infra_mongo_repository "go-microservice-boilerplate-api/infra/repository/mongo"
	"go-microservice-boilerplate-api/infra/secret"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var SecretService = secret.Adapter(secret.CreateSecret())
var PostgresService = infra_database.Adapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())
var MongoService = infra_database.Adapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())

var MongoRepository = infra_repository.Adapter[*core_dog_entity.DogEntity](infra_mongo_repository.CreateMongoRepository[*core_dog_entity.DogEntity]())

func init() {
	SecretService.InitEnvs()
	PostgresService.Connect()
	MongoService.Connect()

	var filter = infra_repository.FindOneInput{}
	result, err := MongoRepository.FindOne(filter.SetMongoFilter(&bson.D{{Key: "name", Value: "Hiquinho"}}))

	fmt.Println(result, err)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.SendString("up and running")
	})

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
