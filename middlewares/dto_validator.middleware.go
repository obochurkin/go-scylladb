package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateDTO(dto interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(dto); err != nil {
			log.Print(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := validate.Struct(dto); err != nil {
			log.Print(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Store the validated DTO in the request context
		c.Set("validatedDTO", dto)

		c.Next()
	}
}
