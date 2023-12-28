package heartrate_v1

import (
	"github.com/gin-gonic/gin"

	"github.com/obochurkin/go-scylladb/middlewares"
)

// SetupRoutes initializes the application routes.
func SetupHeartbeatV1Routes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/heart-rate/:id", GetHeartRateByPetChipID)
		v1.POST("/heart-rate", middlewares.ValidateDTO(&HeartRateV1RequestDto{}), AddHeartRateByPetChipID)
	}
}
