package core_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
)

type CatCreateAdapter interface {
	CatCreateExecute(cat *core_cat_entity.CatEntity) (string, error)
}

func CatCreateUsecase(repository core_cat_repository.CatRepositoryAdapter) func(input *core_cat_entity.CatEntity) (string, error) {
	return func(input *core_cat_entity.CatEntity) (string, error) {
		catEntity, err := repository.Base().Create(input, "cats")
		if err != nil {
			return "", err
		}
		return catEntity, nil
	}
}
