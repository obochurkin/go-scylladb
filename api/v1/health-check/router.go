package healthCheck

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the application routes.
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/health-check", HealthCheckHandler)
	}
}