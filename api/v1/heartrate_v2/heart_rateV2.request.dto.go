package heartrate_v2

import (
	"github.com/gocql/gocql"
)

type HeartRateV2RequestDto struct {
	PetChipID gocql.UUID `json:"petChipId" binding:"required,max=50"`
	HeartRate int        `json:"heartRete" binding:"required,min=0,max=500"`
}
