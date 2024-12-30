package modules_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	infra "go-microservice-boilerplate-api/infra/logger"
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var catRepository core_cat_repository.CatRepositoryAdapter = CreateCatRepository()
var loggerService infra.LoggerAdapter = infra.LoggerAdapter(infra.CreateLogger())

type CatCreateAdapter interface {
	CatCreateExecute(cat *core_cat_entity.CatEntity) (utils.Nullable[string], *utils.AppException)
	Validate(cat *core_cat_entity.CatEntity) error
}

func CatCreate() CatCreateAdapter {
	return &CreateAdatper{}
}

type CatDeleteAdapter interface {
	CatDeleteExecute(id string) *utils.AppException
	Validate(id primitive.ObjectID) error
}

func CatDelete() CatDeleteAdapter {
	return &DeleteAdatper{}
}

type CatGetByIDAdapter interface {
	CatGetByIDExecute(input string) (*core_cat_entity.CatEntity, *utils.AppException)
	Validate(input primitive.ObjectID) error
}

func CatGetByID() CatGetByIDAdapter {
	return &GetByIDAdatper{}
}

type CatUpdateAdapter interface {
	CatUpdateExecute(cat *core_cat_entity.CatEntity) (*core_cat_entity.CatEntity, *utils.AppException)
	Validate(cat *core_cat_entity.CatEntity) utils.Nullable[string]
}

func CatUpdate() CatUpdateAdapter {
	return &UpdateAdatper{}
}

type CatListAdapter interface {
	CatListExecute(input utils.PaginationType) ([]core_cat_entity.CatEntity, *utils.AppException)
}

func CatList() CatListAdapter {
	return &ListAdatper{}
}
