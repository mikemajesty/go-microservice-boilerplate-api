package core_cat

import (
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
)

type CatGetByIDAdapter interface {
	CatGetByIDExecute(cat string) (*core_cat.CatEntity, error)
}

func CatGetByIDUsecase(repository core_cat_repository.CatRepositoryAdapter) func(input string) (*core_cat.CatEntity, error) {
	return func(input string) (*core_cat.CatEntity, error) {
		filter := infra_repository.FindOneInput[string]{}
		entity, err := repository.Base().FindByID(filter.CreatePostgresFilter(&utils.Entity[string]{ID: input}), "cats")

		if err != nil {
			return entity, err
		}
		return entity, nil
	}
}
