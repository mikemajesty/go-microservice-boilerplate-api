package infra_database

import (
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type databaseType interface {
	*mongo.Client | *gorm.DB
}

type DatabaseAdapter[T databaseType] interface {
	Connect() (T, *utils.AppException)
	DB() T
}
