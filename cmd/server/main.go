package main

import (
	"github.com/gin-gonic/gin"

	"github.com/obochurkin/go-scylladb/api/v1/buildings"
	"github.com/obochurkin/go-scylladb/api/v1/health-check"
	"github.com/obochurkin/go-scylladb/api/v1/heartrate_v1"
	"github.com/obochurkin/go-scylladb/api/v2/heartrate_v2"
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
  healthCheck.SetupRoutes(r)
	heartrate_v1.SetupRoutes(r)
  heartrate_v2.SetupRoutes(r)
	buildings.SetupRoutes(r)
	r.Run(":8080")
}
