package core_dog_entity

import (
	utils_entity "go-microservice-boilerplate-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DogEntity struct {
	utils_entity.Entity[*primitive.ObjectID] `bson:",inline"`
	Name                                     string `bson:"name"`
}

func (entity *DogEntity) Build(name string) *DogEntity {
	entity.Name = name
	entity.Entity = utils_entity.Entity[*primitive.ObjectID]{}
	return entity
}
