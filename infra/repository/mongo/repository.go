package infra_mongo_repository

import (
	"context"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/infra/secret"
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDatabase = infra_database.DatabaseAdapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())
var SecretService = secret.SecretAdapter(secret.CreateSecret())

type adapter[T utils.EntityAdapter] struct{}

func (a *adapter[T]) Create(entity T, table string) (string, utils.ApiException) {
	var databse = SecretService.GetSecret("MONGO_INITDB_DATABASE")
	var id = primitive.NewObjectID()
	entity.SetCreatedAt()
	entity.SetUpdatedAt()
	entity.SetID(id)
	result, err := MongoDatabase.DB().Database(databse).Collection(table).InsertOne(context.Background(), entity)
	if err != nil {
		return "", utils.ApiInternalServerException(err.Error())
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *adapter[T]) FindByID(input *infra_repository.FindOneInput[primitive.ObjectID], table string) (T, utils.ApiException) {
	var databse = SecretService.GetSecret("MONGO_INITDB_DATABASE")
	var entity = new(T)
	var err = MongoDatabase.DB().Database(databse).Collection(table).FindOne(context.Background(), input.MongoFilter).Decode(&entity)
	if err != nil {
		return *entity, utils.ApiNotFoundException(err.Error())
	}
	return *entity, nil
}

func (a *adapter[T]) Update(entity T, table string) (string, utils.ApiException) {
	var databse = SecretService.GetSecret("MONGO_INITDB_DATABASE")
	entity.SetUpdatedAt()
	_, err := MongoDatabase.DB().Database(databse).Collection(table).UpdateOne(context.Background(), bson.M{"_id": entity.GetID()}, bson.M{"$set": entity})
	if err != nil {
		return entity.GetID().(primitive.ObjectID).Hex(), utils.ApiNotFoundException(err.Error())
	}
	return entity.GetID().(primitive.ObjectID).Hex(), nil
}

func (a *adapter[T]) Delete(entity T, table string) utils.ApiException {
	var databse = SecretService.GetSecret("MONGO_INITDB_DATABASE")
	_, err := MongoDatabase.DB().Database(databse).Collection(table).DeleteOne(context.Background(), bson.M{"_id": entity.GetID()})
	if err != nil {
		return utils.ApiInternalServerException(err.Error())
	}
	return nil
}

func (a *adapter[T]) List(table string) ([]T, utils.ApiException) {
	var databse = SecretService.GetSecret("MONGO_INITDB_DATABASE")
	var entities []T
	cursor, err := MongoDatabase.DB().Database(databse).Collection(table).Find(context.Background(), bson.M{})
	if err != nil {
		return entities, utils.ApiNotFoundException(err.Error())
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var entity T
		err := cursor.Decode(&entity)
		if err != nil {
			return entities, utils.ApiInternalServerException(err.Error())
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func CreateMongoRepository[T utils.EntityAdapter]() infra_repository.IRepository[T, primitive.ObjectID] {
	return &adapter[T]{}
}
