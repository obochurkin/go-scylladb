package heartrate_v2

import (
	"time"

	"github.com/gocql/gocql"
)

// a few words about TTL it can be used in update and insert operations
// TTL measures in seconds
// after expiration data will be lost
// default TTL value is null, which make data persistent
// ex. UPDATE heartrate USING TTL 600 SET heart_rate = 110 WHERE pet_chip_id = 123e4567-e89b-12d3-a456-426655440b23;
// get TTL of the field ex SELECT name, TTL(heart_rate)FROM heartrate WHERE  pet_chip_id = 123e4567-e89b-12d3-a456-426655440b23;
// applying TTL to a record on ex. INSERT INTO heartrate(pet_chip_id, name, heart_rate) VALUES (c63e71f0-936e-11ea-bb37-0242ac130002, 'Rocky', 87) USING TTL 30;


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