package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
)

type DogListAdapter interface {
	DogListExecute() ([]core_dog.DogEntity, error)
}

func DogListUsecase(repository core_dog_repository.DogRepositoryAdapter) func() ([]core_dog.DogEntity, error) {
	return func() ([]core_dog.DogEntity, error) {
		entityList, err := repository.Base().List("dogs")

		var dogEntities []core_dog.DogEntity

		if err != nil {
			return dogEntities, err
		}

		for _, dog := range entityList {
			dogEntities = append(dogEntities, *dog)
		}
		return dogEntities, nil
	}
}
