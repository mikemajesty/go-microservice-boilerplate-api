package infra_repository

import (
	utils "go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type Adapter[T utils.EntityAdapter, P utils.EntityIDAdapter] interface {
	IRepository[T, P]
}

type FindOneInput[T utils.EntityIDAdapter] struct {
	MongoFilter    *bson.D
	PostgresFilter *utils.Entity[T]
}

func (f FindOneInput[T]) CreateMongoFilter(filter *bson.D) *FindOneInput[T] {
	f.MongoFilter = filter
	return &f
}

func (f FindOneInput[T]) CreatePostgresFilter(filter *utils.Entity[T]) *FindOneInput[T] {
	f.PostgresFilter = filter
	return &f
}

type IRepository[T utils.EntityAdapter, P utils.EntityIDAdapter] interface {
	Create(entity T, table string) (string, *utils.AppException)
	FindByID(input *FindOneInput[P], table string) (T, *utils.AppException)
	Update(entity T, table string) (string, *utils.AppException)
	Delete(entity T, table string) *utils.AppException
	List(table string) ([]T, *utils.AppException)
}
