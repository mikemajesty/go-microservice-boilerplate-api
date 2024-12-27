package core_cat

import (
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
)

type CatDeleteAdapter interface {
	CatDeleteExecute(cat string) error
}

func CatDeleteUsecase(repository core_cat_repository.CatRepositoryAdapter) func(input string) error {
	return func(input string) error {
		filter := infra_repository.FindOneInput[string]{}
		entity, errNotFound := repository.Base().FindByID(filter.CreatePostgresFilter(&utils.Entity[string]{ID: input}), "cats")

		if errNotFound != nil {
			return errNotFound
		}
		err := repository.Base().Delete(entity, "cats")
		if err != nil {
			return err
		}
		return nil
	}
}
