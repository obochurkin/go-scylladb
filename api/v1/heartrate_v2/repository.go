package heartrate_v2

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	InsertHeartRate = "INSERT INTO heartrate_v2 (pet_chip_id, time, heart_rate) VALUES (?,toUnixTimestamp(now()),?);"
	SelectHeartRateByPetChipIDAndTimeRange = "SELECT pet_chip_id, time, heart_rate FROM heartrate_v2 WHERE pet_chip_id = ? AND time >= ? AND time <= ?;"
	SelectLastAddedRecordByPetChipID = "SELECT pet_chip_id, time, heart_rate FROM heartrate_v2 WHERE pet_chip_id = ? ORDER BY time DESC LIMIT 1;"
)

func InsertHeartRateQuery(session *gocql.Session, petChipID gocql.UUID, heartRate int) error {
	return session.Query(InsertHeartRate, petChipID, heartRate).Exec()
}

func SelectHeartRateByChipQuery(session *gocql.Session, petChipID gocql.UUID, startTimeRange *time.Time, endTimeRange *time.Time) ([]HeartRateV2, error) {
	var heartRates []HeartRateV2

	iter := session.Query(SelectHeartRateByPetChipIDAndTimeRange, petChipID, startTimeRange.UTC(), endTimeRange.UTC()).Iter()
	defer iter.Close()

	for {
		var heartRate HeartRateV2
		if !iter.Scan(&heartRate.ID, &heartRate.Time, &heartRate.HeartRate) {
			break
		}
		heartRate.Time = heartRate.Time.Local()
		heartRates = append(heartRates, heartRate)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return heartRates, nil
}

func SelectLastAddedRecordByPetChipIDQuery(session *gocql.Session, petChipID gocql.UUID) (HeartRateV2, error) {
	var heartRate HeartRateV2
	if err := session.Query(SelectLastAddedRecordByPetChipID, petChipID).Scan(&heartRate.ID, &heartRate.Time, &heartRate.HeartRate); err != nil {
			return HeartRateV2{}, err
	}
	return heartRate, nil
}