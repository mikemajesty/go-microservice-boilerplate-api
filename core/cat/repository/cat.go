package core_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
)

type CatRepositoryAdapter interface {
	Base() infra_repository.IRepository[*core_cat_entity.CatEntity, string]
	Paginate(input utils.ListInput) ([]core_cat_entity.CatEntity, *utils.AppException)
}
