package buildings

import (
	"github.com/gin-gonic/gin"

	"github.com/obochurkin/go-scylladb/middlewares"
)

// SetupRoutes initializes the application routes.
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.POST("/buildings", middlewares.ValidateDTO(&AddBuildingDto{}), AddBuilding)
		v1.GET("/buildings", middlewares.DynamicQueryValidator(func() interface{} { return &GetBuildingQueryParam{} }), GetBuilding)
	}
}