package core_cat

import (
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	"go-microservice-boilerplate-api/utils"
)

func CatListUsecase(repository core_cat_repository.CatRepositoryAdapter) func(input utils.PostgresListInput) ([]core_cat.CatEntity, *utils.AppException) {
	return func(input utils.PostgresListInput) ([]core_cat.CatEntity, *utils.AppException) {
		entityList, err := repository.Paginate(utils.PostgresListInput{Pagination: input.Pagination, Sort: input.Sort})

		var catEntities []core_cat.CatEntity

		if err != nil {
			return catEntities, err
		}

		catEntities = append(catEntities, entityList...)
		return catEntities, nil
	}
}
