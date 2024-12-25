package core_dog_entity

import (
	utils_entity "go-microservice-boilerplate-api/utils"
)

type DogEntity struct {
	*utils_entity.Entity `bson:",inline"`
	Name                 string `bson:"name"`
}

func (d *DogEntity) Build(name string) DogEntity {
	d.Name = name
	d.Entity = &utils_entity.Entity{}
	return *d
}
