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

var PostgresDatabase = infra_database.Adapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())

type adapter[T utils_entity.IEntity] struct{}

func (a *adapter[T]) Create(entity T) (string, error) {
	entity.SetID(primitive.NewObjectID().Hex())
	entity.SetCreatedAt()
	entity.SetUpdatedAt()
	fmt.Println(entity.GetID())
	result := PostgresDatabase.DB().Create(&entity)
	if result.Error != nil {
		return entity.GetID().(string), result.Error

	}
	return entity.GetID().(string), nil
}

func (a *adapter[T]) FindOne(input *infra_repository.FindOneInput[string]) (T, error) {
	fmt.Println(input.PostgresFilter, "input")
	var entity = new(T)
	err := PostgresDatabase.DB().First(entity, input.PostgresFilter)
	if err.Error != nil {
		return *entity, fmt.Errorf("failed to find one: %s", err.Error.Error())
	}
	return *entity, nil
}

func CreatePostgresRepository[T utils_entity.IEntity]() infra_repository.IRepository[T, string] {
	return &adapter[T]{}
}
