package core_dog_repository

import core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"

type Adapter interface {
	FindOne(filter *core_dog_entity.DogEntity) (core_dog_entity.DogEntity, error)
	Create(data core_dog_entity.DogEntity) (string, error)
}
