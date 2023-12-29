package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DynamicQueryValidator(newInstanceFunc func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParamsDto := newInstanceFunc()
		// Bind the query parameters to the struct
		if err := c.ShouldBindQuery(queryParamsDto); err != nil {
			log.Printf("Validation error: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		log.Print(queryParamsDto)
		// Validate the struct fields
		if err := validate.StructPartial(queryParamsDto); err != nil {
			log.Printf("Validation error: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Store the validated query parameters in the context if needed
		c.Set("query", queryParamsDto)

		// Continue to the next middleware or handler
		c.Next()
	}
}
