package healthCheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// HealthCheckHandler is the handler for the health check endpoint.
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}