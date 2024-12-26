package core_dog_repository

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	infra_repository "go-microservice-boilerplate-api/infra/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDogRepository interface {
	Base() infra_repository.IRepository[*core_dog_entity.DogEntity, *primitive.ObjectID]
}
