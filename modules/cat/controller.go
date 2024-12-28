package modules_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"

	"github.com/gin-gonic/gin"
)

func CreateCat(controller *gin.Context) {
	var entity = core_cat_entity.CatEntity{}
	controller.Bind(&entity)

	result, err := CatCreate().CatCreateExecute(&entity)

	if err != nil {
		controller.JSON(err.GetStatus(), err.GetMessage())
		return
	}

	controller.JSON(200, result)
}

func DeleteCat(controller *gin.Context) {
	id := controller.Param("id")
	err := CatDelete().CatDeleteExecute(id)

	if err != nil {
		controller.JSON(err.GetStatus(), err.Message)
		return
	}

	controller.JSON(200, "Success")
}

func GetCatByID(controller *gin.Context) {
	id := controller.Param("id")
	result, err := CatGetByID().CatGetByIDExecute(id)

	if err != nil {
		controller.JSON(err.GetStatus(), err.Message)
		return
	}

	controller.JSON(200, result)
}

func UpdateCat(controller *gin.Context) {
	var entity = core_cat_entity.CatEntity{}

	entity.SetID(controller.Param("id"))
	controller.Bind(&entity)

	result, err := CatUpdate().CatUpdateExecute(&entity)

	if err != nil {
		controller.JSON(err.GetStatus(), err.Message)
		return
	}

	controller.JSON(200, result)
}

func ListCat(controller *gin.Context) {
	result, err := CatList().CatListExecute()

	if err != nil {
		controller.JSON(err.GetStatus(), err.Message)
		return
	}

	controller.JSON(200, result)
}
