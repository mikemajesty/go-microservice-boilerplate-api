package core_cat

import (
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
)

type CatListAdapter interface {
	CatListExecute() ([]core_cat.CatEntity, error)
}

func CatListUsecase(repository core_cat_repository.CatRepositoryAdapter) func() ([]core_cat.CatEntity, error) {
	return func() ([]core_cat.CatEntity, error) {
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