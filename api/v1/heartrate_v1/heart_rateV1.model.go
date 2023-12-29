package heartrate_v1

import (
	"time"

	"github.com/gocql/gocql"
)

type HeartRateV1 struct {
	ID        gocql.UUID `json:"pet_chip_id"`
	Time      time.Time `json:"time"`
	HeartRate int       `json:"heart_rate"`
}
