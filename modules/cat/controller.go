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
		controller.JSON(500, err.Error())
		return
	}

	controller.JSON(200, result)
}

func DeleteCat(controller *gin.Context) {
	id := controller.Param("id")
	err := CatDelete().CatDeleteExecute(id)

	if err != nil {
		controller.JSON(404, err.Error())
		return
	}

	controller.JSON(200, "Success")
}
