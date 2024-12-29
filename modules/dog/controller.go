package modules_dog

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	infra "go-microservice-boilerplate-api/infra/logger"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var LoggerService = infra.LoggerAdapter(infra.CreateLogger())

func CreateDog(controller *gin.Context) {
	var entity = core_dog_entity.DogEntity{}
	controller.Bind(&entity)
	result, err := DogCreate().DogCreateExecute(&entity)

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, result)
}

func DeleteDog(controller *gin.Context) {
	id := controller.Param("id")
	err := DogDelete().DogDeleteExecute(id)

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, "Success")
}

func GetDog(controller *gin.Context) {
	id := controller.Param("id")
	result, err := DogGetByID().DogGetByIDExecute(id)

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, result)
}

func ListDog(controller *gin.Context) {
	result, err := DogList().DogListExecute()

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
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
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, result)
}
