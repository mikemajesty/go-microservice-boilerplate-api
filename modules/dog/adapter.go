package modules_dog

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	"go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dogRepository core_dog_repository.DogRepositoryAdapter = CreateDogRepository()

type DogCreateAdapter interface {
	DogCreateExecute(dog *core_dog_entity.DogEntity) (utils.Nullable[string], *utils.AppException)
	Validate(dog *core_dog_entity.DogEntity) error
}

func DogCreate() DogCreateAdapter {
	return &CreateAdatper{}
}

type DogDeleteAdapter interface {
	DogDeleteExecute(id string) *utils.AppException
	Validate(input primitive.ObjectID) error
}

func DogDelete() DogDeleteAdapter {
	return &DeleteAdatper{}
}

type DogGetByIDAdapter interface {
	DogGetByIDExecute(id string) (*core_dog_entity.DogEntity, *utils.AppException)
	Validate(input primitive.ObjectID) error
}

func DogGetByID() DogGetByIDAdapter {
	return &GetByIDAdatper{}
}

type DogListAdapter interface {
	DogListExecute(input utils.MongoListInput) ([]core_dog_entity.DogEntity, *utils.AppException)
}

func DogList() DogListAdapter {
	return &ListAdatper{}
}

type DogUpdateAdapter interface {
	DogUpdateExecute(dog *core_dog_entity.DogEntity) (*core_dog_entity.DogEntity, *utils.AppException)
	Validate(dog *core_dog_entity.DogEntity) utils.Nullable[string]
}

func DogUpdate() DogUpdateAdapter {
	return &UpdateAdatper{}
}
