package core_cat

import (
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	"go-microservice-boilerplate-api/utils"
)

type CatListAdapter interface {
	CatListExecute() ([]core_cat.CatEntity, *utils.AppException)
}

func CatListUsecase(repository core_cat_repository.CatRepositoryAdapter) func() ([]core_cat.CatEntity, *utils.AppException) {
	return func() ([]core_cat.CatEntity, *utils.AppException) {
		entityList, err := repository.Base().List("cats")

		var catEntities []core_cat.CatEntity

		if err != nil {
			return catEntities, err
		}

		for _, cat := range entityList {
			catEntities = append(catEntities, *cat)
		}
		return catEntities, nil
	}
}
