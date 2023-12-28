package main

import (
	"github.com/gin-gonic/gin"

	"github.com/obochurkin/go-scylladb/api/v1/health-check"
	"github.com/obochurkin/go-scylladb/api/v1/heartrate_v1"
	"github.com/obochurkin/go-scylladb/internal"
	"github.com/obochurkin/go-scylladb/middlewares"
)

func main() {
	//init gin
	r := gin.Default()
	r.Use(middlewares.ErrorHandlerMiddleware)

	//init DB
	internal.InitDB()

	session := internal.GetSession()
	defer session.Close()

	//init routes
  healthCheck.SetupHealthCheckRoutes(r)
	heartrate_v1.SetupHeartbeatV1Routes(r)
	r.Run(":8080")
}
