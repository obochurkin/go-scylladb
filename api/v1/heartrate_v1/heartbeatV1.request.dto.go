package heartrate_v1

import (
	"github.com/gocql/gocql"
)

type HeartRateV1RequestDto struct {
	PetChipID gocql.UUID `json:"petChipId" binding:"required"`
	HeartRate int       `json:"heartRete" binding:"required,min=0,max=500"`
}
