package heartrate_v2

import (
	"time"

	"github.com/gocql/gocql"
)

type HeartRateV2 struct {
	ID        gocql.UUID `json:"pet_chip_id"`
	Time      time.Time `json:"time"`
	HeartRate int       `json:"heart_rate"`
}
