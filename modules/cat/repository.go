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

func (r *repository) Paginate(input utils.PostgresListInput) ([]core_cat_entity.CatEntity, *utils.AppException) {
	skip := int(input.Pagination.Page-1) * input.Pagination.Limit
	limit := input.Pagination.Limit
	var cats []core_cat_entity.CatEntity

	result := MongoDatabase.DB().Limit(limit).Offset(skip).Order(input.Sort).Find(&cats, GetFilter(input))

	if result.Error != nil {
		return nil, utils.ApiInternalServerException(result.Error.Error())
	}
	return cats, nil
}

func CreateCatRepository() core_cat_repository.CatRepositoryAdapter {
	return &repository{}
}

func (r *repository) Base() infra_repository.IRepository[*core_cat_entity.CatEntity, string] {
	return CatRepository
}

func GetFilter(input utils.PostgresListInput) utils.Nullable[map[string]interface{}] {
	if input.Search.Field == "" {
		return interface{}(nil)
	}
	return map[string]interface{}{input.Search.Field: input.Search.Value}
}
