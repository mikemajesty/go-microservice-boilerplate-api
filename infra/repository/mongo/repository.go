package infra_mongo_repository

import (
	"context"
	"fmt"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	utils_entity "go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDatabase = infra_database.Adapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())

type adapter[T utils_entity.IEntity] struct{}

func (a *adapter[T]) Create(entity T) (string, error) {
	entity.SetID(primitive.NewObjectID())
	entity.SetCreatedAt()
	entity.SetUpdatedAt()
	result, err := MongoDatabase.DB().Database("go-microservice-boilerplate-api").Collection("dog").InsertOne(context.Background(), entity)
	if err != nil {
		return "", fmt.Errorf("failed to create: %s", err.Error())
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *adapter[T]) FindOne(input *infra_repository.FindOneInput[*primitive.ObjectID]) (T, error) {
	var entity = new(T)
	var err = MongoDatabase.DB().Database("go-microservice-boilerplate-api").Collection("dog").FindOne(context.Background(), input.MongoFilter).Decode(&entity)
	if err != nil {
		return *entity, fmt.Errorf("failed to find one: %s", err.Error())
	}
	return *entity, nil
}

func CreateMongoRepository[T utils_entity.IEntity]() infra_repository.IRepository[T, *primitive.ObjectID] {
	return &adapter[T]{}
}
