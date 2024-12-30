package modules_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	infra_postgres_repository "go-microservice-boilerplate-api/infra/repository/postgres"
	"go-microservice-boilerplate-api/utils"
)

var CatRepository = infra_repository.Adapter[*core_cat_entity.CatEntity, string](infra_postgres_repository.CreatePostgresRepository[*core_cat_entity.CatEntity]())
var MongoDatabase = infra_database_postgres.CreateConnectPostgres()

type repository struct{}

func (r *repository) Paginate(input utils.ListInput) ([]core_cat_entity.CatEntity, *utils.AppException) {
	skip := int(input.Pagination.Page-1) * input.Pagination.Limit
	limit := input.Pagination.Limit
	var cats []core_cat_entity.CatEntity
	result := MongoDatabase.DB().Limit(limit).Offset(skip).Find(&cats)

	if result.Error != nil {
		return nil, utils.ApiInternalServerException(result.Error.Error())
	}
	return cats, nil
}

func CreateCatRepository() core_cat_repository.CatRepositoryAdapter {
	return &repository{}
}

func (r *repository) FindByName(input string) (*core_cat_entity.CatEntity, error) {
	var entity = core_cat_entity.CatEntity{Name: input}
	result := MongoDatabase.DB().Find(&entity, "name = ?", input)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (r *repository) Base() infra_repository.IRepository[*core_cat_entity.CatEntity, string] {
	return CatRepository
}
