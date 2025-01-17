package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DogGetByIDUsecase(repository core_dog_repository.DogRepositoryAdapter) func(id string) (*core_dog.DogEntity, *utils.AppException) {
	return func(id string) (*core_dog.DogEntity, *utils.AppException) {
		filter := infra_repository.FindOneInput[primitive.ObjectID]{}
		objectID, _ := primitive.ObjectIDFromHex(id)

		entity, err := repository.Base().FindByID(filter.CreateMongoFilter(&bson.D{{Key: "_id", Value: objectID}}), "dogs")

		if err != nil {
			return entity, err
		}
		return entity, nil
	}
}
