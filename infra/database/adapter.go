package infra_database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type databaseType interface {
	*mongo.Client | *gorm.DB
}

type DatabaseAdapter[T databaseType] interface {
	Connect() (T, error)
	DB() T
}
