package modules_cat

import (
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	core_usecase_cat "go-microservice-boilerplate-api/core/cat/use-case"
	infra "go-microservice-boilerplate-api/infra/logger"
	"go-microservice-boilerplate-api/utils"
)

var catRepository core_cat_repository.CatRepositoryAdapter = CreateCatRepository()
var loggerService infra.LoggerAdapter = infra.LoggerAdapter(infra.CreateLogger())

type createAdatper struct{}

func CatCreate() core_usecase_cat.CatCreateAdapter {
	return &createAdatper{}
}

func (c *createAdatper) CatCreateExecute(cat *core_cat_entity.CatEntity) (utils.Nullable[string], *utils.AppException) {
	return core_usecase_cat.CatCreateUsecase(catRepository, loggerService)(cat)
}

type deleteAdatper struct{}

func (d *deleteAdatper) CatDeleteExecute(input string) *utils.AppException {
	return core_usecase_cat.CatDeleteUsecase(catRepository)(input)
}

func CatDelete() core_usecase_cat.CatDeleteAdapter {
	return &deleteAdatper{}
}

type getByIDAdatper struct{}

func (d *getByIDAdatper) CatGetByIDExecute(input string) (*core_cat.CatEntity, *utils.AppException) {
	return core_usecase_cat.CatGetByIDUsecase(catRepository)(input)
}

func CatGetByID() core_usecase_cat.CatGetByIDAdapter {
	return &getByIDAdatper{}
}

type updateAdatper struct{}

func (d *updateAdatper) CatUpdateExecute(input *core_cat.CatEntity) (*core_cat.CatEntity, *utils.AppException) {
	return core_usecase_cat.CatUpdateUsecase(catRepository)(input)
}

func CatUpdate() core_usecase_cat.CatUpdateAdapter {
	return &updateAdatper{}
}

type listAdatper struct{}

func (d *listAdatper) CatListExecute() ([]core_cat.CatEntity, *utils.AppException) {
	return core_usecase_cat.CatListUsecase(catRepository)()
}

func CatList() core_usecase_cat.CatListAdapter {
	return &listAdatper{}
}
