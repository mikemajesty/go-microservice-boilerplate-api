package main

import (
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	"go-microservice-boilerplate-api/infra/secret"
	modules_cat "go-microservice-boilerplate-api/modules/cat"
	modules_dog "go-microservice-boilerplate-api/modules/dog"
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var SecretService = secret.Adapter(secret.CreateSecret())
var PostgresService = infra_database.Adapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())
var MongoService = infra_database.Adapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())

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
	utils.Route.GET("/", func(c *gin.Context) {
		c.JSON(200, "up and running")
	})
	modules_cat.CreateCatRoutes()
	modules_dog.CreateDogRoutes()

	utils.Route.Run()
}
