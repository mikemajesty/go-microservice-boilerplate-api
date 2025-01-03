package modules_cat

import (
	core_cat_entity "go-microservice-boilerplate-api/core/cat/entity"
	infra "go-microservice-boilerplate-api/infra/logger"
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
)

var LoggerService = infra.LoggerAdapter(infra.CreateLogger())

func CreateCat(controller *gin.Context) {
	var entity = core_cat_entity.CatEntity{}
	controller.Bind(&entity)

	result, err := CatCreate().CatCreateExecute(&entity)

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, result)
}

func DeleteCat(controller *gin.Context) {
	id := controller.Param("id")
	err := CatDelete().CatDeleteExecute(id)

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, "Success")
}

func GetCatByID(controller *gin.Context) {
	id := controller.Param("id")
	result, err := CatGetByID().CatGetByIDExecute(id)

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
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
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, result)
}

func ListCat(controller *gin.Context) {
	sort, err := utils.PostgresSort(controller)

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	search, errSearch := utils.CreateSearch(controller)

	if errSearch != nil {
		LoggerService.Error(errSearch.GetMessage(), infra.LogAttrInput{"status": errSearch.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(errSearch.GetStatus(), errSearch.Response(errSearch.GetStatus(), trace.(string)))
		return
	}

	pagination := utils.Pagination(controller)

	result, err := CatList().CatListExecute(utils.PostgresListInput{Sort: sort.(string), Pagination: pagination, Search: search})

	if err != nil {
		LoggerService.Error(err.GetMessage(), infra.LogAttrInput{"status": err.GetStatus()})
		trace, _ := controller.Get("traceId")
		controller.JSON(err.GetStatus(), err.Response(err.GetStatus(), trace.(string)))
		return
	}

	controller.JSON(200, result)
}
