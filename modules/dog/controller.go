package modules_dog

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"

	"github.com/gin-gonic/gin"
)

var dogRepository = CreateDogRepository()

func CreateDog(controller *gin.Context) {
	var entity = core_dog_entity.DogEntity{}
	controller.Bind(&entity)
	result, err := dogRepository.Base().Create(&entity, "dog")

	if err != nil {
		controller.JSON(500, err.Error())
		return
	}

	controller.JSON(200, result)
}
