package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	"go-microservice-boilerplate-api/utils"
)

type DogCreateAdapter interface {
	DogCreateExecute(dog *core_dog.DogEntity) (string, *utils.AppException)
}

func DogCreateUsecase(repository core_dog_repository.DogRepositoryAdapter) func(input *core_dog.DogEntity) (string, *utils.AppException) {
	return func(input *core_dog.DogEntity) (string, *utils.AppException) {
		dogEntity, err := repository.Base().Create(input, "dogs")
		if err != nil {
			return "", err
		}
		return dogEntity, nil
	}
}
