package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	"go-microservice-boilerplate-api/utils"
)

func DogListUsecase(repository core_dog_repository.DogRepositoryAdapter) func() ([]core_dog.DogEntity, *utils.AppException) {
	return func() ([]core_dog.DogEntity, *utils.AppException) {
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
