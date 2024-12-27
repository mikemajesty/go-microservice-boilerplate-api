package core_dog

import (
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DogDeleteAdapter interface {
	DogDeleteExecute(id string) error
}

func DogDeleteUsecase(repository core_dog_repository.DogRepositoryAdapter) func(id string) error {
	return func(id string) error {
		filter := infra_repository.FindOneInput[primitive.ObjectID]{}
		objectID, _ := primitive.ObjectIDFromHex(id)

		entity, errNotFound := repository.Base().FindByID(filter.CreateMongoFilter(&bson.D{{Key: "_id", Value: objectID}}), "dogs")

		if errNotFound != nil {
			return errNotFound
		}

		err := repository.Base().Delete(entity, "dogs")
		if err != nil {
			return err
		}
		return nil
	}
}