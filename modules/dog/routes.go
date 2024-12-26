package modules_dog

import (
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateDogRoutes() *gin.Engine {
	utils.Route.POST("/dogs", CreateDog)
	return utils.Route
}
