package heartrate_v1

import (
	"time"

	"github.com/gocql/gocql"
)

type HeartRateV1ResponseDto struct {
	ID gocql.UUID `json:"petChipID"`
	Time time.Time `json:"time"`
	HeartRate int `json:"heartRate"`
}

func newHeartRateV1ResponseDto(model *HeartRateV1) *HeartRateV1ResponseDto {
	return &HeartRateV1ResponseDto{
		ID: model.ID,
		HeartRate: model.HeartRate,
		Time: model.Time,
	}
}