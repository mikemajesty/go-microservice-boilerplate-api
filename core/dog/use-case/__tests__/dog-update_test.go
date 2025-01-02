package core_dog

import (
	entity "go-microservice-boilerplate-api/core/dog/entity"
	repository "go-microservice-boilerplate-api/core/dog/repository"
	usecase "go-microservice-boilerplate-api/core/dog/use-case"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_When_Dog_Update_NotFound_Should_Expect_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.DogRepositoryAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.DogEntity, primitive.ObjectID]]()

	When(repo.Base()).ThenReturn(repoBase)

	findByIDIinput := infra_repository.FindOneInput[primitive.ObjectID]{}
	objectID, _ := primitive.ObjectIDFromHex("6775d3762b2ff3103c6de582")
	When(repoBase.FindByID(findByIDIinput.CreateMongoFilter(&primitive.D{{Key: "_id", Value: objectID}}), "dogs")).ThenReturn(nil, utils.ApiNotFoundException("Dog not found"))

	id, _ := primitive.ObjectIDFromHex("6775d3762b2ff3103c6de582")
	dog := entity.DogEntity{Name: "Marley", Entity: utils.Entity[primitive.ObjectID]{ID: id}}
	_, err := usecase.DogUpdateUsecase(repo)(&dog)

	if err == nil {
		t.Fail()
	}

	if err != nil {
		assert.Equal(t, err.Status, 404)
	}
}

func Test_When_Dog_Update_Successfully_Should_Expect_No_Error(t *testing.T) {
	SetUp(t)
	repo := Mock[repository.DogRepositoryAdapter]()
	repoBase := Mock[infra_repository.IRepository[*entity.DogEntity, primitive.ObjectID]]()

	When(repo.Base()).ThenReturn(repoBase)

	findByIDIinput := infra_repository.FindOneInput[primitive.ObjectID]{}
	objectID, _ := primitive.ObjectIDFromHex("6775d3762b2ff3103c6de582")
	id, _ := primitive.ObjectIDFromHex("6775d3762b2ff3103c6de582")
	dog := entity.DogEntity{Name: "Marley", Entity: utils.Entity[primitive.ObjectID]{ID: id}}
	When(repoBase.FindByID(findByIDIinput.CreateMongoFilter(&primitive.D{{Key: "_id", Value: objectID}}), "dogs")).ThenReturn(&dog, nil)
	When(repoBase.Update(&dog, "dogs")).ThenReturn("6775d3762b2ff3103c6de582", nil)

	entity, err := usecase.DogUpdateUsecase(repo)(&dog)

	if err != nil {
		t.Fail()
	}

	assert.NotNil(t, entity)

}
