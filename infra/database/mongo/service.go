package infra_database_mongo

import (
	"context"
	"fmt"
	infra_database "go-microservice-boilerplate-api/infra/database"
	"go-microservice-boilerplate-api/infra/secret"
	"go-microservice-boilerplate-api/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var EnvService = secret.SecretAdapter(secret.CreateSecret())

type adapter struct{}

func CreateConnectMongo() infra_database.DatabaseAdapter[*mongo.Client] {
	return &adapter{}
}

var db *mongo.Client

func (adapter *adapter) DB() *mongo.Client {
	return db
}

func (adapter *adapter) Connect() (*mongo.Client, utils.ApiException) {
	uri := EnvService.GetSecret("MONGO_URI")
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	db = client

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, utils.ApiInternalServerException(err.Error())
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}
