package core_dog_entity

import (
	utils_entity "go-microservice-boilerplate-api/utils"
)

type DogEntity struct {
	utils_entity.Entity `bson:",inline"`
	Name                string `bson:"name"`
}

func (entity *DogEntity) Build(name string) *DogEntity {
	entity.Name = name
	entity.Entity = utils_entity.Entity{}
	return entity
}
