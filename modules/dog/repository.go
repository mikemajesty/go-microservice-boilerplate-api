package modules_dog

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_repository "go-microservice-boilerplate-api/core/dog/repository"
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	infra_mongo_repository "go-microservice-boilerplate-api/infra/repository/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DogRepository = infra_repository.Adapter[*core_dog_entity.DogEntity, *primitive.ObjectID](infra_mongo_repository.CreateMongoRepository[*core_dog_entity.DogEntity]())

type repository struct{}

func CreateDogRepository() core_dog_repository.IDogRepository {
	return &repository{}
}

func (r *repository) Base() infra_repository.IRepository[*core_dog_entity.DogEntity, *primitive.ObjectID] {
	return DogRepository
}
