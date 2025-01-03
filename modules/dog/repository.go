package modules_dog

import (
	"context"
	"fmt"
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_mongo "go-microservice-boilerplate-api/infra/database/mongo"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	infra_mongo_repository "go-microservice-boilerplate-api/infra/repository/mongo"
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DogRepository = infra_repository.Adapter[*core_dog_entity.DogEntity, primitive.ObjectID](infra_mongo_repository.CreateMongoRepository[*core_dog_entity.DogEntity]())
var PostgresService = infra_database.DatabaseAdapter[*mongo.Client](infra_database_mongo.CreateConnectMongo())

type repository struct{}

func (r *repository) Paginate(input utils.MongoListInput) ([]core_dog_entity.DogEntity, *utils.AppException) {
	skip := int64(input.Pagination.Page-1) * int64(input.Pagination.Limit)
	limit := int64(input.Pagination.Limit)

	fOpt := options.FindOptions{Limit: &limit, Skip: &skip}

	_context := context.Background()

	cursor, err := PostgresService.DB().Database("go-microservice-boilerplate-api").Collection("dogs").Find(_context, GetFilter(input), fOpt.SetSort(CreateMongoSort(input.Sort)))

	var entities []core_dog_entity.DogEntity

	if err != nil {
		return entities, utils.ApiNotFoundException(err.Error())
	}

	defer cursor.Close(_context)

	fmt.Println(cursor, "cursor")
	for cursor.Next(_context) {
		var entity core_dog_entity.DogEntity
		err := cursor.Decode(&entity)
		if err != nil {
			return entities, utils.ApiInternalServerException(err.Error())
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func CreateMongoSort(sort []utils.MongoSortType) bson.D {
	var sortFields bson.D
	for _, s := range sort {
		sortFields = append(sortFields, bson.E{Key: s.Field, Value: s.Order})
	}

	return sortFields
}

func CreateDogRepository() core_dog_repository.DogRepositoryAdapter {
	return &repository{}
}

func (r *repository) Base() infra_repository.IRepository[*core_dog_entity.DogEntity, primitive.ObjectID] {
	return DogRepository
}

func GetFilter(input utils.MongoListInput) bson.D {
	if input.Search.Field == "" {
		return bson.D{}

	}
	filter := bson.D{{Key: input.Search.Field, Value: input.Search.Value}}
	return filter
}
