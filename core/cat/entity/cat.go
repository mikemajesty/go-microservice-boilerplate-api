package core_cat

import (
	utils_entity "go-microservice-boilerplate-api/utils"
)

type CatEntity struct {
	utils_entity.Entity[string] `gorm:"embedded"`
	Name                        string `bson:"name" gorm:"column:name"`
}

func (entity *CatEntity) Build(name string) *CatEntity {
	entity.Name = name
	entity.Entity = utils_entity.Entity[string]{}
	return entity
}

func (entity *CatEntity) TableName() string {
	return "cats"
}
