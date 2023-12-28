package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"

	users "github.com/obochurkin/go-scylladb/api/v1/users/models"
	"github.com/obochurkin/go-scylladb/internal"
)

// HealthCheckHandler is the handler for the health check endpoint.
func GetUserHandler(ctx *gin.Context) {
	var user users.User

	userID := ctx.Param("id")
	// Check if the ID is a valid integer
	if _, err := strconv.Atoi(userID); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	session := internal.GetSession()

	err := session.Query("SELECT user_id, fname, lname FROM users WHERE user_id = ?;", userID).WithContext(ctx).Scan(&user.ID, &user.FirstName, &user.LastName)

	if err == gocql.ErrNotFound {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"firstName": &user.FirstName,
		"lastName":  &user.LastName,
	})
}
