package modules_dog

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	core_dog_usecase "go-microservice-boilerplate-api/core/dog/use-case"
	"go-microservice-boilerplate-api/utils"
)

var dogRepository core_dog_repository.DogRepositoryAdapter = CreateDogRepository()

type createAdatper struct{}

func DogCreate() core_dog_usecase.DogCreateAdapter {
	return &createAdatper{}
}

func (c *createAdatper) DogCreateExecute(dog *core_dog_entity.DogEntity) (string, *utils.AppException) {
	return core_dog_usecase.DogCreateUsecase(dogRepository)(dog)
}

type deleteAdatper struct{}

func DogDelete() core_dog_usecase.DogDeleteAdapter {
	return &deleteAdatper{}
}

func (c *deleteAdatper) DogDeleteExecute(id string) *utils.AppException {
	return core_dog_usecase.DogDeleteUsecase(dogRepository)(id)
}

type getByIDAdatper struct{}

func DogGetByID() core_dog_usecase.DogGetByIDAdapter {
	return &getByIDAdatper{}
}

func (c *getByIDAdatper) DogGetByIDExecute(id string) (*core_dog_entity.DogEntity, *utils.AppException) {
	return core_dog_usecase.DogGetByIDUsecase(dogRepository)(id)
}

type listAdatper struct{}

func DogList() core_dog_usecase.DogListAdapter {
	return &listAdatper{}
}

func (c *listAdatper) DogListExecute() ([]core_dog_entity.DogEntity, *utils.AppException) {
	return core_dog_usecase.DogListUsecase(dogRepository)()
}

type updateAdatper struct{}

func DogUpdate() core_dog_usecase.DogUpdateAdapter {
	return &updateAdatper{}
}

func (c *updateAdatper) DogUpdateExecute(dog *core_dog_entity.DogEntity) (*core_dog_entity.DogEntity, *utils.AppException) {
	return core_dog_usecase.DogUpdateUsecase(dogRepository)(dog)
}
