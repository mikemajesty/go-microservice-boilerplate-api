package infra_repository

import (
	utils_entity "go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type Adapter[T utils_entity.IEntity] interface {
	IRepository[T]
}

type FindOneInput struct {
	MongoFilter    *bson.D
	PostgresFilter *utils_entity.Entity
}

func (f FindOneInput) SetMongoFilter(filter *bson.D) *FindOneInput {
	f.MongoFilter = filter
	return &f
}

func (f FindOneInput) SetPostgresFilter(filter *utils_entity.Entity) *FindOneInput {
	f.PostgresFilter = filter
	return &f
}

type IRepository[T utils_entity.IEntity] interface {
	Create(entity T) (string, error)
	FindOne(input *FindOneInput) (T, error)
	// Update(entity T) error
	// FindByID(id int) (*T, error)
	// FindAll() ([]*T, error)
	// DeleteByID(id int) error
}
