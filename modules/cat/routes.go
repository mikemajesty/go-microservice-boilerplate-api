package modules_cat

import (
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateCatRoutes() *gin.Engine {
	utils.Route.POST("/cats", CreateCat)
	return utils.Route
}
