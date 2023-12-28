package heartrate_v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"

	"github.com/obochurkin/go-scylladb/internal"
	"github.com/obochurkin/go-scylladb/utils"
)

func GetHeartRateByPetChipID(ctx *gin.Context) {
	id := ctx.Param("id")

	uuid, validationError := utils.IsValidUUID(id)
	if validationError != nil {
		log.Printf("Invalid UUID: %v", validationError)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	session := internal.GetSession()

	heartRate, err := SelectHeartRateByChipQuery(session, uuid)
	if err != nil {
		log.Printf("Query error: %v", err)
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"petChipID": heartRate.ID,
		"time":      heartRate.Time,
		"heartRate": heartRate.HeartRate,
	})
}

func AddHeartRateByPetChipID(ctx *gin.Context) {
	// Retrieve the validated DTO from the context
	dtoInterface, exists := ctx.Get("validatedDTO")
	if !exists {
		log.Print("DTO not found in context")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Assert the DTO to the specific type
	dto, ok := dtoInterface.(*HeartRateV1RequestDto)
	if !ok {
		log.Print("Invalid DTO type in context")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	session := internal.GetSession()

	// Insert the heart rate
	if err := InsertHeartRateQuery(session, dto.PetChipID, dto.HeartRate); err != nil {
		log.Printf("Insertion error: %v", err)
		handleError(ctx, err)
		return
	}

	// Retrieve the inserted heart rate
	heartRate, err := SelectHeartRateByChipQuery(session, dto.PetChipID)
	if err != nil {
		log.Printf("Query error: %v", err)
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"petChipID": heartRate.ID,
		"time":      heartRate.Time,
		"heartRate": heartRate.HeartRate,
	})
}

func handleError(ctx *gin.Context, err error) {
	switch err {
	case gocql.ErrNotFound:
		ctx.AbortWithStatus(http.StatusNotFound)
	default:
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
