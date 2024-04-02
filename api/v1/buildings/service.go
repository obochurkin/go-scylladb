package buildings

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/obochurkin/go-scylladb/internal"
)

type AddBuildingDto struct {
	Name         string `json:"name" binding:"required,max=50"`
	City         string `json:"city" binding:"required,max=50"`
	BuiltYear    int    `json:"built_year" binding:"required,max=5000"`
	HeightMeters int    `json:"height_meters" binding:"required,max=10000000"`
}

type GetBuildingQueryParam struct {
	City *string `form:"city" validate:"omitempty,max=50"`
}

func AddBuilding(ctx *gin.Context) {
	// Retrieve the validated DTO from the context
	dtoInterface, exists := ctx.Get("validatedDTO")
	if !exists {
		log.Print("DTO not found in context")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Assert the DTO to the specific type
	dto, ok := dtoInterface.(*AddBuildingDto)
	if !ok {
		log.Print("Invalid DTO type in context")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	session := internal.GetSession()

	if err := InsertBuildingQuery(session, dto); err != nil {
		log.Printf("Insertion error: %v", err)
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "null",
	})
}

func GetBuilding(ctx *gin.Context) {
	queryInterface, _ := ctx.Get("query")
	query, ok := queryInterface.(*GetBuildingQueryParam)
	if !ok {
		log.Print("Invalid Query type in context")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Set("query", nil)

	session := internal.GetSession() 

	data, err := GetBuildingByCity(session, query)
	if err != nil {
		log.Printf("Query error: %v", err)
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
