package infra_postgres_repository

import (
	"fmt"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

var PostgresDatabase = infra_database.DatabaseAdapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())

type adapter[T utils.EntityAdapter] struct{}

func (a *adapter[T]) Create(entity T, table string) (string, utils.ApiException) {
	entity.SetID(primitive.NewObjectID().Hex())
	entity.SetCreatedAt()
	entity.SetUpdatedAt()
	fmt.Println(entity.GetID())
	result := PostgresDatabase.DB().Table(table).Create(&entity)
	if result.Error != nil {
		return entity.GetID().(string), utils.ApiInternalServerException(result.Error.Error())

	}
	return entity.GetID().(string), nil
}

func (a *adapter[T]) FindByID(input *infra_repository.FindOneInput[string], table string) (T, utils.ApiException) {
	fmt.Println(input.PostgresFilter, "input")
	var entity = new(T)
	err := PostgresDatabase.DB().Table(table).First(entity, input.PostgresFilter)
	if err.Error != nil {
		return *entity, utils.ApiNotFoundException(err.Error.Error())
	}
	return *entity, nil
}

func (a *adapter[T]) Update(entity T, table string) (string, utils.ApiException) {
	entity.SetUpdatedAt()
	result := PostgresDatabase.DB().Table(table).Save(&entity)
	if result.Error != nil {
		return entity.GetID().(string), utils.ApiNotFoundException(result.Error.Error())
	}
	return entity.GetID().(string), nil
}

func (a *adapter[T]) Delete(entity T, table string) utils.ApiException {
	fmt.Println(entity.GetID())
	result := PostgresDatabase.DB().Table(table).Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *adapter[T]) List(table string) ([]T, utils.ApiException) {
	var entities []T
	result := PostgresDatabase.DB().Table(table).Find(&entities)
	if result.Error != nil {
		return nil, utils.ApiNotFoundException(result.Error.Error())
	}
	return entities, nil
}

func CreatePostgresRepository[T utils.EntityAdapter]() infra_repository.IRepository[T, string] {
	return &adapter[T]{}
}
