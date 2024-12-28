package infra_repository

import (
	"go-microservice-boilerplate-api/utils"
	utils_entity "go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type Adapter[T utils_entity.EntityAdapter, P utils_entity.EntityIDAdapter] interface {
	IRepository[T, P]
}

type FindOneInput[T utils_entity.EntityIDAdapter] struct {
	MongoFilter    *bson.D
	PostgresFilter *utils_entity.Entity[T]
}

func (f FindOneInput[T]) CreateMongoFilter(filter *bson.D) *FindOneInput[T] {
	f.MongoFilter = filter
	return &f
}

func (f FindOneInput[T]) CreatePostgresFilter(filter *utils_entity.Entity[T]) *FindOneInput[T] {
	f.PostgresFilter = filter
	return &f
}

type IRepository[T utils_entity.EntityAdapter, P utils_entity.EntityIDAdapter] interface {
	Create(entity T, table string) (string, utils.ApiException)
	FindByID(input *FindOneInput[P], table string) (T, utils.ApiException)
	Update(entity T, table string) (string, utils.ApiException)
	Delete(entity T, table string) utils.ApiException
	List(table string) ([]T, utils.ApiException)
}
