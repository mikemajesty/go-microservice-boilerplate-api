package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DogUpdateAdapter interface {
	DogUpdateExecute(dog *core_dog.DogEntity) (*core_dog.DogEntity, *utils.AppException)
}

func DogUpdateUsecase(repository core_dog_repository.DogRepositoryAdapter) func(input *core_dog.DogEntity) (*core_dog.DogEntity, *utils.AppException) {
	return func(input *core_dog.DogEntity) (*core_dog.DogEntity, *utils.AppException) {
		errorList := validate(input)

		if len(errorList) > 0 {
			return input, utils.ApiBadRequestException(errorList)
		}

		filter := infra_repository.FindOneInput[primitive.ObjectID]{}

		entity, err := repository.Base().FindByID(filter.CreateMongoFilter(&bson.D{{Key: "_id", Value: input.GetID()}}), "dogs")

		if err != nil {
			return entity, utils.ApiNotFoundException("dog with id " + input.ConvertIDToString() + " not found")
		}

		entity = input

		repository.Base().Update(entity, "dogs")

		return entity, nil
	}
}

func validate(input *core_dog.DogEntity) string {
	objectIdError := utils.IsObjectID(input.ID, "ID")

	errorList := []string{}

	if objectIdError != nil {
		errorList = append(errorList, objectIdError.(error).Error())
	}

	nameEror := utils.ValidateSchema(input.Name, "Name", validation.Required, is.LowerCase, is.Alpha, validation.Length(3, 50))

	if nameEror != nil {
		errorList = append(errorList, nameEror.Error())
	}

	return strings.Join(errorList, ",")
}
