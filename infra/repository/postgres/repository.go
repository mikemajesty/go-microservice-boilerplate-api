package infra_postgres_repository

import (
	"fmt"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	utils_entity "go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

var PostgresDatabase = infra_database.DatabaseAdapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())

type adapter[T utils_entity.EntityAdapter] struct{}

func (a *adapter[T]) Create(entity T, table string) (string, error) {
	entity.SetID(primitive.NewObjectID().Hex())
	entity.SetCreatedAt()
	entity.SetUpdatedAt()
	fmt.Println(entity.GetID())
	result := PostgresDatabase.DB().Table(table).Create(&entity)
	if result.Error != nil {
		return entity.GetID().(string), result.Error

	}
	return entity.GetID().(string), nil
}

func (a *adapter[T]) FindByID(input *infra_repository.FindOneInput[string], table string) (T, error) {
	fmt.Println(input.PostgresFilter, "input")
	var entity = new(T)
	err := PostgresDatabase.DB().Table(table).First(entity, input.PostgresFilter)
	if err.Error != nil {
		return *entity, fmt.Errorf("failed to find one: %s", err.Error.Error())
	}
	return *entity, nil
}

func (a *adapter[T]) Update(entity T, table string) (string, error) {
	entity.SetUpdatedAt()
	result := PostgresDatabase.DB().Table(table).Save(&entity)
	if result.Error != nil {
		return entity.GetID().(string), fmt.Errorf("failed to update: %s", result.Error)
	}
	return entity.GetID().(string), nil
}

func (a *adapter[T]) Delete(entity T, table string) error {
	fmt.Println(entity.GetID())
	result := PostgresDatabase.DB().Table(table).Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *adapter[T]) List(table string) ([]T, error) {
	var entities []T
	result := PostgresDatabase.DB().Table(table).Find(&entities)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list: %s", result.Error)
	}
	return entities, nil
}

func CreatePostgresRepository[T utils_entity.EntityAdapter]() infra_repository.IRepository[T, string] {
	return &adapter[T]{}
}
