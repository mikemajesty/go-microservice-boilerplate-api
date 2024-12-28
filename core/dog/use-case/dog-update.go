package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DogUpdateAdapter interface {
	DogUpdateExecute(dog *core_dog.DogEntity) (*core_dog.DogEntity, *utils.AppException)
}

func DogUpdateUsecase(repository core_dog_repository.DogRepositoryAdapter) func(input *core_dog.DogEntity) (*core_dog.DogEntity, *utils.AppException) {
	return func(input *core_dog.DogEntity) (*core_dog.DogEntity, *utils.AppException) {
		filter := infra_repository.FindOneInput[primitive.ObjectID]{}

		entity, err := repository.Base().FindByID(filter.CreateMongoFilter(&bson.D{{Key: "_id", Value: input.GetID()}}), "dogs")

		if err != nil {
			return entity, err
		}

		entity = input

		repository.Base().Update(entity, "dogs")

		return entity, nil
	}
}
