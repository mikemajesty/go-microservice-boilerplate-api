package main

import (
	"fmt"
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	infra_mongo_repository "go-microservice-boilerplate-api/infra/repository/mongo"
	infra_postgres_repository "go-microservice-boilerplate-api/infra/repository/postgres"
	"go-microservice-boilerplate-api/infra/secret"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var SecretService = secret.Adapter(secret.CreateSecret())
var PostgresService = infra_database.Adapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())
var MongoService = infra_database.Adapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())

var DogRepository = core_dog_repository.DogRepository(infra_mongo_repository.CreateMongoRepository[*core_dog_entity.DogEntity]())
var CatRepository = core_cat_repository.CatRepository(infra_postgres_repository.CreatePostgresRepository[*core_cat_entity.CatEntity]())

func init() {
	SecretService.InitEnvs()
	PostgresService.Connect()
	MongoService.Connect()
}

// func list(){

// 	mongoList, errMongo := MongoRepository.List("dog")

// 	if errMongo != nil {
// 		log.Fatal(errMongo)
// 	}

// 	for _, v := range mongoList {
// 		fmt.Println(v, "mongo")
// 	}

// 	postgresList, errPostgres := PostgresRepository.List("cats")

// 	if errPostgres != nil {
// 		log.Fatal(errPostgres)
// 	}

// 	for _, v := range postgresList {
// 		fmt.Println(v, "postgres")
// 	}
// }

// func deletePostgres() {
// 	var filter = infra_repository.FindOneInput[string]{}

// 	result, err := PostgresRepository.FindByID(filter.CreatePostgresFilter(&utils_entity.Entity[string]{ID: "676c8e26b2fcc097c9897776"}), "cats")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	PostgresRepository.Delete(result, "cats")

// }

// func deteteMongo() {
// objectID, _ := primitive.ObjectIDFromHex("676c5eabcec5a5a7eddd8f59")

// var filter = infra_repository.FindOneInput[*primitive.ObjectID]{}

// var result, err = MongoRepository.FindByID(filter.CreateMongoFilter(&bson.D{{Key: "_id", Value: objectID}}), "dog")

// if err != nil {
// 	log.Fatal(err)
// }

// MongoRepository.Delete(result, "dog")
// }

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.SendString("up and running")
	})

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
