package core_cat

import (
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
)

func CatDeleteUsecase(repository core_cat_repository.CatRepositoryAdapter) func(id string) *utils.AppException {
	return func(id string) *utils.AppException {
		filter := infra_repository.FindOneInput[string]{}
		entity, errNotFound := repository.Base().FindByID(filter.CreatePostgresFilter(&utils.Entity[string]{ID: id}), "cats")

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
