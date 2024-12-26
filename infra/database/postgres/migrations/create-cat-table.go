package main

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	infra_database "go-microservice-boilerplate-api/infra/database"
	infra_database_postgres "go-microservice-boilerplate-api/infra/database/postgres"
	"go-microservice-boilerplate-api/infra/secret"

	"gorm.io/gorm"
)

var SecretService = secret.Adapter(secret.CreateSecret())
var PostgresService = infra_database.Adapter[*gorm.DB](infra_database_postgres.CreateConnectPostgres())

func init() {
	SecretService.InitEnvs()
	PostgresService.Connect()
}

func main() {
	PostgresService.DB().AutoMigrate(&core_cat_entity.CatEntity{})
}
