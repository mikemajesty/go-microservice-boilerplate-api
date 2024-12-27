package modules_cat

import (
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateCatRoutes() *gin.Engine {
	utils.Route.POST("/cats", CreateCat)
	utils.Route.DELETE("/cats/:id", DeleteCat)
	utils.Route.GET("/cats/:id", GetCatByID)
	utils.Route.PUT("/cats/:id", UpdateCat)
	utils.Route.GET("/cats", ListCat)
	return utils.Route
}
