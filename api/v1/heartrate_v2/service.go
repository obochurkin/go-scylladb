package heartrate_v2

import (
	"log"
	"net/http"
	"time"

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

	// Retrieve the validated DTO from the context
	queryInterface, exists := ctx.Get("query")
	if !exists {
		log.Print("Query not found in context")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Assert the DTO to the specific type
	query, ok := queryInterface.(*HeartRateV2QueryParams)
	if !ok {
		log.Print("Invalid Query type in context")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Set("query", nil)

	// set default end date to time now
	if query.EndDate.IsZero() {
		log.Print("Set default end date to time now...")
		query.EndDate = time.Now()
	}

	if status := validateTimeRange(ctx, query.StartDate.Format(time.RFC3339), query.EndDate.Format(time.RFC3339)); !status {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	session := internal.GetSession()

	heartRates, err := SelectHeartRateByChipQuery(session, uuid, &query.StartDate, &query.EndDate)
	if err != nil {
		log.Printf("Query error: %v", err)
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK,  heartRates)
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
	dto, ok := dtoInterface.(*HeartRateV2RequestDto)
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
	heartRate, err := SelectLastAddedRecordByPetChipIDQuery(session, dto.PetChipID)
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

func validateDateTime(ctx *gin.Context, dateTime string) bool {
	time, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		log.Printf("Invalid RFC3339 datetime: %v", err)
		return false
	}
	if time.IsZero() {
		return false
	}
	return true
}

func validateTimeRange(ctx *gin.Context, startDateStr, endDateStr string) bool {
	// Validate start date and end date individually
	startValid := validateDateTime(ctx, startDateStr)
	endValid := validateDateTime(ctx, endDateStr)

	if !startValid || !endValid {
		return false
	}

	// Convert the strings to time.Time objects
	startDate, err1 := time.Parse(time.RFC3339, startDateStr)
	endDate, err2 := time.Parse(time.RFC3339, endDateStr)

	if err1 != nil || err2 != nil {
		log.Printf("Error parsing dates: %v, %v", err1, err2)
		return false
	}

	// Check if startDate is less than or equal to endDate
	if startDate.After(endDate) {
		log.Println("Start date cannot be after end date")
		return false
	}

	return true
}
