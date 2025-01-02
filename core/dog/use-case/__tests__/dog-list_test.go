package core_dog

import (
	entity "go-microservice-boilerplate-api/core/dog/entity"
	repository "go-microservice-boilerplate-api/core/dog/repository"
	usecase "go-microservice-boilerplate-api/core/dog/use-case"
	"go-microservice-boilerplate-api/utils"
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/assert"
)

// Test_When_Dog_Paginate_Successfully_Should_Expect_List
func Test_When_Dog_Paginate_Error_Should_Expect_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.DogRepositoryAdapter]()

	list := []entity.DogEntity{}
	When(repo.Paginate(Any[utils.MongoListInput]())).ThenReturn(list, utils.ApiInternalServerException("Error"))

	_, err := usecase.DogListUsecase(repo)(utils.MongoListInput{})

	if err == nil {
		t.Fail()
	}

	if err != nil {
		assert.Equal(t, err.Status, 500)
	}
}

func Test_When_Dog_Paginate_Successfully_Should_Expect_List(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.DogRepositoryAdapter]()

	list := []entity.DogEntity{
		{Name: "Marley"},
		{Name: "Marley"},
	}
	When(repo.Paginate(Any[utils.MongoListInput]())).ThenReturn(list, nil)

	result, err := usecase.DogListUsecase(repo)(utils.MongoListInput{})

	if err != nil {
		t.Fail()
	}

	if err == nil {
		assert.Equal(t, result, list)
	}
}
