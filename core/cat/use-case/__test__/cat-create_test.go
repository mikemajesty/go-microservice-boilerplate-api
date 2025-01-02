package core_dog

import (
	entity "go-microservice-boilerplate-api/core/cat/entity"
	repository "go-microservice-boilerplate-api/core/cat/repository"
	usecase "go-microservice-boilerplate-api/core/cat/use-case"
	infra "go-microservice-boilerplate-api/infra/logger"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
)

func Test_When_Cat_Create_Successfully_Should_Expect_ID(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.CatRepositoryAdapter]()
	logger := Mock[infra.LoggerAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.CatEntity, string]]()
	mock := entity.CatEntity{
		Name: "Marley",
	}

	When(repo.Base()).ThenReturn(repoBase)
	When(repoBase.Create(&mock, "cats")).ThenReturn("6775d3762b2ff3103c6de582", nil)

	result, err := usecase.CatCreateUsecase(repo, logger)(&mock)

	if result != "6775d3762b2ff3103c6de582" {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}
