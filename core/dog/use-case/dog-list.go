package core_dog

import (
	core_dog "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	"go-microservice-boilerplate-api/utils"
)

func DogListUsecase(repository core_dog_repository.DogRepositoryAdapter) func(input utils.MongoListInput) ([]core_dog.DogEntity, *utils.AppException) {
	return func(input utils.MongoListInput) ([]core_dog.DogEntity, *utils.AppException) {
		entityList, err := repository.Paginate(utils.MongoListInput{Pagination: input.Pagination, Sort: input.Sort})

		var dogEntities []core_dog.DogEntity

		if err != nil {
			return dogEntities, err
		}

		dogEntities = append(dogEntities, entityList...)

		return dogEntities, nil
	}
}
