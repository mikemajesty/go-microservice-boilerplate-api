package core_cat

import (
	"fmt"
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	"go-microservice-boilerplate-api/utils"
)

func CatListUsecase(repository core_cat_repository.CatRepositoryAdapter) func(input utils.PaginationType) ([]core_cat.CatEntity, *utils.AppException) {
	return func(input utils.PaginationType) ([]core_cat.CatEntity, *utils.AppException) {
		fmt.Println("CatListUsecase", input.Limit)
		fmt.Println("CatListUsecase", input.Page)
		entityList, err := repository.Paginate(utils.ListInput{Pagination: input})

		var catEntities []core_cat.CatEntity

		if err != nil {
			return catEntities, err
		}

		catEntities = append(catEntities, entityList...)
		return catEntities, nil
	}
}
