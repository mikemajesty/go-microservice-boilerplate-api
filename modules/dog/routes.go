package modules_dog

import (
	"go-microservice-boilerplate-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateDogRoutes() *gin.Engine {
	utils.Route.POST("/dogs", CreateDog)
	utils.Route.DELETE("/dogs/:id", DeleteDog)
	utils.Route.GET("/dogs/:id", GetDog)
	utils.Route.GET("/dogs", ListDog)
	utils.Route.PUT("/dogs/:id", UpdateDog)
	return utils.Route
}
