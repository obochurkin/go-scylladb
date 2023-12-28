package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next() // Continue processing

	status := c.Writer.Status()

	if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
		switch status {
		case http.StatusNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		case http.StatusBadRequest:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		case http.StatusUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		case http.StatusForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		case http.StatusMethodNotAllowed:
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
		default:
			c.JSON(status, gin.H{"error": http.StatusText(status)})
		}
		return
	}

	if status >= http.StatusInternalServerError {
		log.Println("Internal Server Error:", c.Errors.Last())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
}
