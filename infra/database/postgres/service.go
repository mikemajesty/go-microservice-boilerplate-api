package infra_database_postgres

import (
	"fmt"
	infra_database "go-microservice-boilerplate-api/infra/database"
	"go-microservice-boilerplate-api/infra/secret"
	"go-microservice-boilerplate-api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var EnvService = secret.SecretAdapter(secret.CreateSecret())

type adapter struct{}

func CreateConnectPostgres() infra_database.DatabaseAdapter[*gorm.DB] {
	return &adapter{}
}

var db *gorm.DB

func (adapter *adapter) DB() *gorm.DB {
	return db
}

func (adapter *adapter) Connect() (*gorm.DB, utils.ApiException) {
	host := EnvService.GetSecret("POSTGRES_HOST")
	port := EnvService.GetSecret("POSTGRES_PORT")
	user := EnvService.GetSecret("POSTGRES_USER")
	pass := EnvService.GetSecret("POSTGRES_PASSWORD")
	database := EnvService.GetSecret("POSTGRES_DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, database, port)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, utils.ApiInternalServerException(err.Error())
	}
	db = con

	fmt.Println("Successfully connected to Postgres")
	return con, nil
}
