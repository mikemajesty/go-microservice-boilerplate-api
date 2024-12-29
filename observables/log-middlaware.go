package observables

import (
	infra "go-microservice-boilerplate-api/infra/logger"

	"github.com/gin-gonic/gin"
)

func GinBodyLogMiddleware(logger infra.LoggerAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		statusCode := c.Writer.Status()
		if statusCode < 400 {
			logger.Info("request completed", infra.LogAttrInput{"method": c.Request.Method, "status": statusCode, "path": c.Request.URL.Path})
		}
	}
}
