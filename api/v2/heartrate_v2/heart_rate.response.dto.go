package heartrate_v2

import (
	"time"

	"github.com/gocql/gocql"
)

type HeartRateV2ResponseDto struct {
	ID gocql.UUID `json:"petChipID"`
	Time time.Time `json:"time"`
	HeartRate int `json:"heartRate"`
}

func newHeartRateV2ResponseDto(model *HeartRateV2) *HeartRateV2ResponseDto {
	return &HeartRateV2ResponseDto{
		ID: model.ID,
		HeartRate: model.HeartRate,
		Time: model.Time,
	}
}