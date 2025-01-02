package core_dog

import (
	entity "go-microservice-boilerplate-api/core/cat/entity"
	repository "go-microservice-boilerplate-api/core/cat/repository"
	usecase "go-microservice-boilerplate-api/core/cat/use-case"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/assert"
)

func Test_When_Cat_GetByID_NotFound_Should_Expect_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.CatRepositoryAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.CatEntity, string]]()

	When(repo.Base()).ThenReturn(repoBase)

	findByIDIinput := infra_repository.FindOneInput[string]{}
	When(repoBase.FindByID(findByIDIinput.CreatePostgresFilter(&utils.Entity[string]{ID: "6775d3762b2ff3103c6de582"}), "cats")).ThenReturn(nil, utils.ApiNotFoundException("Cat not found"))

	result, err := usecase.CatGetByIDUsecase(repo)("6775d3762b2ff3103c6de582")

	if err == nil {
		t.Fail()
	}

	if err != nil {
		assert.Nil(t, result)
	}
}

func Test_When_Cat_GetByID_Successfully_Should_Expect_No_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.CatRepositoryAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.CatEntity, string]]()

	When(repo.Base()).ThenReturn(repoBase)

	findByIDIinput := infra_repository.FindOneInput[string]{}

	cat := entity.CatEntity{Name: "Marley"}
	When(repoBase.FindByID(findByIDIinput.CreatePostgresFilter(&utils.Entity[string]{ID: "6775d3762b2ff3103c6de582"}), "cats")).ThenReturn(&cat, nil)

	entity, err := usecase.CatGetByIDUsecase(repo)("6775d3762b2ff3103c6de582")

	if err != nil {
		t.Fail()
	}

	if err == nil {
		assert.Equal(t, entity, &cat)
	}
}
