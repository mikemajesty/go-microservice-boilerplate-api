package modules_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	core_cat_repository "go-microservice-boilerplate-api/core/cat/repository"
	core_usecase_cat "go-microservice-boilerplate-api/core/cat/use-case"
)

var catRepository core_cat_repository.CatRepositoryAdapter = CreateCatRepository()

type createAdatper struct{}

func CatCreate() core_usecase_cat.CatCreateAdapter {
	return &createAdatper{}
}

func (c *createAdatper) CatCreateExecute(cat *core_cat_entity.CatEntity) (string, error) {
	return core_usecase_cat.CatCreateUsecase(catRepository)(cat)
}
