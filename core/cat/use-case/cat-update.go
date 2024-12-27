package core_cat

import (
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
)

type CatUpdateAdapter interface {
	CatUpdateExecute(cat *core_cat.CatEntity) (*core_cat.CatEntity, error)
}

func CatUpdateUsecase(repository core_cat_repository.CatRepositoryAdapter) func(input *core_cat.CatEntity) (*core_cat.CatEntity, error) {
	return func(input *core_cat.CatEntity) (*core_cat.CatEntity, error) {
		filter := infra_repository.FindOneInput[string]{}
		entity, err := repository.Base().FindByID(filter.CreatePostgresFilter(&utils.Entity[string]{ID: input.ID}), "cats")

		if err != nil {
			return entity, err
		}

		entity = input

		repository.Base().Update(entity, "cats")

		return entity, nil
	}
}