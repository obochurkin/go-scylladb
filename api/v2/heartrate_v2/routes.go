package heartrate_v2

import (
	"github.com/gin-gonic/gin"

	"github.com/obochurkin/go-scylladb/middlewares"
)

// SetupRoutes initializes the application routes.
func SetupRoutes(router *gin.Engine) {
	v2 := router.Group("api/v2")
	{
		v2.GET("/heart-rate/:id", middlewares.DynamicQueryValidator(func() interface{} { return &HeartRateV2QueryParams{} }), GetHeartRateByPetChipID)
		v2.POST("/heart-rate", middlewares.ValidateDTO(&HeartRateV2RequestDto{}), AddHeartRateByPetChipID)
	}
}
