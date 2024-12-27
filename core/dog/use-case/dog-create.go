package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
)

type DogCreateAdapter interface {
	DogCreateExecute(dog *core_dog.DogEntity) (string, error)
}

func DogCreateUsecase(repository core_dog_repository.DogRepositoryAdapter) func(input *core_dog.DogEntity) (string, error) {
	return func(input *core_dog.DogEntity) (string, error) {
		dogEntity, err := repository.Base().Create(input, "dogs")
		if err != nil {
			return "", err
		}
		return dogEntity, nil
	}
}
