package core_dog

import (
	entity "go-microservice-boilerplate-api/core/cat/entity"
	repository "go-microservice-boilerplate-api/core/cat/repository"
	usecase "go-microservice-boilerplate-api/core/cat/use-case"
	"go-microservice-boilerplate-api/utils"
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/assert"
)

func Test_When_Cat_Paginate_Error_Should_Expect_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.CatRepositoryAdapter]()

	list := []entity.CatEntity{}
	When(repo.Paginate(Any[utils.PostgresListInput]())).ThenReturn(list, utils.ApiInternalServerException("Error"))

	_, err := usecase.CatListUsecase(repo)(utils.PostgresListInput{})

	if err == nil {
		t.Fail()
	}

	if err != nil {
		assert.Equal(t, err.Status, 500)
	}
}

func Test_When_Cat_Paginate_Successfully_Should_Expect_List(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.CatRepositoryAdapter]()

	list := []entity.CatEntity{
		{Name: "Marley"},
		{Name: "Marley"},
	}
	When(repo.Paginate(Any[utils.PostgresListInput]())).ThenReturn(list, nil)

	result, err := usecase.CatListUsecase(repo)(utils.PostgresListInput{})

	if err != nil {
		t.Fail()
	}

	if err == nil {
		assert.Equal(t, result, list)
	}
}
