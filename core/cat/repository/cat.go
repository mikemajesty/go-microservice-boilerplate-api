package core_cat_repository

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
)

type CatRepository interface {
	infra_repository.IRepository[*core_cat_entity.CatEntity, string]
}
