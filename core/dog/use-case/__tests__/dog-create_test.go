package core_dog

import (
	entity "go-microservice-boilerplate-api/core/dog/entity"
	repository "go-microservice-boilerplate-api/core/dog/repository"
	usecase "go-microservice-boilerplate-api/core/dog/use-case"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_When_Dog_Create_Successfully_Should_Expect_ID(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.DogRepositoryAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.DogEntity, primitive.ObjectID]]()
	mock := entity.DogEntity{
		Name: "Marley",
	}

	When(repo.Base()).ThenReturn(repoBase)
	When(repoBase.Create(&mock, "dogs")).ThenReturn("6775d3762b2ff3103c6de582", nil)

	result, err := usecase.DogCreateUsecase(repo)(&mock)

	if result != "6775d3762b2ff3103c6de582" {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}
