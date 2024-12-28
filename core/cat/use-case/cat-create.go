package core_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra "go-microservice-boilerplate-api/infra/logger"
	"go-microservice-boilerplate-api/utils"
)

type CatCreateAdapter interface {
	CatCreateExecute(cat *core_cat_entity.CatEntity) (utils.Nullable[string], *utils.AppException)
}

func CatCreateUsecase(repository core_cat_repository.CatRepositoryAdapter, logger infra.LoggerAdapter) func(input *core_cat_entity.CatEntity) (utils.Nullable[string], *utils.AppException) {
	return func(input *core_cat_entity.CatEntity) (utils.Nullable[string], *utils.AppException) {
		catID, err := repository.Base().Create(input, "cats")
		if err != nil {
			return nil, err
		}
		logger.Info("Cat created", infra.InfoAttr{Key: "id", Value: catID})
		return catID, nil
	}
}
