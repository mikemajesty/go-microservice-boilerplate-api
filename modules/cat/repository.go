package modules_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	infra_postgres_repository "go-microservice-boilerplate-api/infra/repository/postgres"
)

var CatRepository = infra_repository.Adapter[*core_cat_entity.CatEntity, string](infra_postgres_repository.CreatePostgresRepository[*core_cat_entity.CatEntity]())

type repository struct{}

func CreateCatRepository() core_cat_repository.ICatRepository {
	return &repository{}
}

func (r *repository) FindByName(entity string) (*core_cat_entity.CatEntity, error) {
	return &core_cat_entity.CatEntity{Name: "MIKE"}, nil
}

func (r *repository) Base() infra_repository.IRepository[*core_cat_entity.CatEntity, string] {
	return CatRepository
}
