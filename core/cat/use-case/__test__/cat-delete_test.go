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

func Test_When_Cat_Delete_NotFound_Should_Expect_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.CatRepositoryAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.CatEntity, string]]()

	When(repo.Base()).ThenReturn(repoBase)

	findByIDIinput := infra_repository.FindOneInput[string]{}

	filter := utils.Entity[string]{ID: "6775d3762b2ff3103c6de582"}

	When(repoBase.FindByID(findByIDIinput.CreatePostgresFilter(&filter), "cats")).ThenReturn(nil, utils.ApiNotFoundException("Cat not found"))

	err := usecase.CatDeleteUsecase(repo)("6775d3762b2ff3103c6de582")

	if err == nil {
		t.Fail()
	}

	if err != nil {
		assert.Equal(t, err.Status, 404)
	}
}

func Test_When_Cat_Delete_Successfully_Should_Expect_No_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.CatRepositoryAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.CatEntity, string]]()

	When(repo.Base()).ThenReturn(repoBase)

	findByIDIinput := infra_repository.FindOneInput[string]{}
	dog := entity.CatEntity{Name: "Marley"}

	filter := utils.Entity[string]{ID: "6775d3762b2ff3103c6de582"}

	When(repoBase.FindByID(findByIDIinput.CreatePostgresFilter(&filter), "dogs")).ThenReturn(&dog, nil)
	When(repoBase.Delete(&dog, "dogs")).ThenReturn(nil)

	err := usecase.CatDeleteUsecase(repo)("6775d3762b2ff3103c6de582")

	if err != nil {
		t.Fail()
	}

	if err == nil {
		assert.Nil(t, err)
	}
}
