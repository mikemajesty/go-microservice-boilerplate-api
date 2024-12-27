package modules_dog

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDog(controller *gin.Context) {
	var entity = core_dog_entity.DogEntity{}
	controller.Bind(&entity)
	result, err := DogCreate().DogCreateExecute(&entity)

	if err != nil {
		controller.JSON(500, err.Error())
		return
	}

	controller.JSON(200, result)
}

func DeleteDog(controller *gin.Context) {
	id := controller.Param("id")
	err := DogDelete().DogDeleteExecute(id)

	if err != nil {
		controller.JSON(500, err.Error())
		return
	}

	controller.JSON(200, "Success")
}

func GetDog(controller *gin.Context) {
	id := controller.Param("id")
	result, err := DogGetByID().DogGetByIDExecute(id)

	if err != nil {
		controller.JSON(500, err.Error())
		return
	}

	controller.JSON(200, result)
}

func ListDog(controller *gin.Context) {
	result, err := DogList().DogListExecute()

	if err != nil {
		controller.JSON(500, err.Error())
		return
	}

	controller.JSON(200, result)
}

func UpdateDog(controller *gin.Context) {
	var entity = core_dog_entity.DogEntity{}
	objectID, _ := primitive.ObjectIDFromHex(controller.Param("id"))
	entity.SetID(objectID)
	controller.Bind(&entity)
	result, err := DogUpdate().DogUpdateExecute(&entity)

	if err != nil {
		controller.JSON(500, err.Error())
		return
	}

	controller.JSON(200, result)
}
