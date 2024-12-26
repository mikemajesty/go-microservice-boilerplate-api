package infra_repository

import (
	utils_entity "go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type Adapter[T utils_entity.IEntity, P utils_entity.IEntityID] interface {
	IRepository[T, P]
}

type FindOneInput[T utils_entity.IEntityID] struct {
	MongoFilter    *bson.D
	PostgresFilter *utils_entity.Entity[T]
}

func (f FindOneInput[T]) SetMongoFilter(filter *bson.D) *FindOneInput[T] {
	f.MongoFilter = filter
	return &f
}

func (f FindOneInput[T]) SetPostgresFilter(filter *utils_entity.Entity[T]) *FindOneInput[T] {
	f.PostgresFilter = filter
	return &f
}

type IRepository[T utils_entity.IEntity, P utils_entity.IEntityID] interface {
	Create(entity T) (string, error)
	FindOne(input *FindOneInput[P]) (T, error)
	// Update(entity T) error
	// FindByID(id int) (*T, error)
	// FindAll() ([]*T, error)
	// DeleteByID(id int) error
}
