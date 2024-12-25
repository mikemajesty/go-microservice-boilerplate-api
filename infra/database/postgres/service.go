package infra_database_postgres

import (
	"errors"
	"fmt"
	infra_database "go-microservice-boilerplate-api/infra/database"
	"go-microservice-boilerplate-api/infra/secret"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var EnvService = secret.Adapter(secret.CreateSecret())

type adapter struct{}

func CreateConnectPostgres() infra_database.Adapter[*gorm.DB] {
	return &adapter{}
}

func (adapter *adapter) DB() *gorm.DB {
	con, err := adapter.Connect()

	if err != nil {
		fmt.Println(err)
	}

	return con
}

func (adapter *adapter) Connect() (*gorm.DB, error) {
	host := EnvService.GetSecret("POSTGRES_HOST")
	port := EnvService.GetSecret("POSTGRES_PORT")
	user := EnvService.GetSecret("POSTGRES_USER")
	pass := EnvService.GetSecret("POSTGRES_PASSWORD")
	database := EnvService.GetSecret("POSTGRES_DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, database, port)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New("failed to connect to Postgres")
	}

	fmt.Println("Successfully connected to Postgres")
	return con, nil
}
