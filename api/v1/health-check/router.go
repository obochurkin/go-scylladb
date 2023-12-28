package healthCheck

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the application routes.
func SetupHealthCheckRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/health-check/:id", HealthCheckHandler)
	}
}