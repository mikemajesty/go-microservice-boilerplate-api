package main

import (
	"fmt"
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	infra_mongo_repository "go-microservice-boilerplate-api/infra/repository/mongo"
	infra_postgres_repository "go-microservice-boilerplate-api/infra/repository/postgres"
	"go-microservice-boilerplate-api/infra/secret"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var SecretService = secret.Adapter(secret.CreateSecret())
var PostgresService = infra_database.Adapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())
var MongoService = infra_database.Adapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())

var MongoRepository = infra_repository.Adapter[*core_dog_entity.DogEntity, *primitive.ObjectID](infra_mongo_repository.CreateMongoRepository[*core_dog_entity.DogEntity]())
var PostgresRepository = infra_repository.Adapter[*core_cat_entity.CatEntity, string](infra_postgres_repository.CreatePostgresRepository[*core_cat_entity.CatEntity]())

func init() {
	SecretService.InitEnvs()
	PostgresService.Connect()
	MongoService.Connect()

	// var filter = infra_repository.FindOneInput[string]{}

	// result, err := PostgresRepository.FindByID(filter.CreatePostgresFilter(&utils_entity.Entity[string]{ID: "676c8e26b2fcc097c9897776"}), "cats")

	// objectID, _ := primitive.ObjectIDFromHex("676c5eabcec5a5a7eddd8f59")

	// var filter = infra_repository.FindOneInput[*primitive.ObjectID]{}

	// var result, err = MongoRepository.FindByID(filter.CreateMongoFilter(&bson.D{{Key: "_id", Value: objectID}}), "dog")

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.SendString("up and running")
	})

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
