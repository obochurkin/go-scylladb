package heartrate_v1

import "github.com/gocql/gocql"

const (
	InsertHeartRate = "INSERT INTO heartrate_v1(pet_chip_id, time, heart_rate) VALUES (?,toUnixTimestamp(now()),?);"
	SelectHeartRateByPetChipID = "SELECT pet_chip_id, time, heart_rate FROM heartrate_v1 WHERE pet_chip_id = ?;"
)

func InsertHeartRateQuery(session *gocql.Session, petChipID gocql.UUID, heartRate int) error {
	return session.Query(InsertHeartRate, petChipID, heartRate).Exec()
}

func SelectHeartRateByChipQuery(session *gocql.Session, petChipID gocql.UUID) (HeartRateV1, error) {
	var heartRate HeartRateV1
	if err := session.Query(SelectHeartRateByPetChipID, petChipID).Scan(&heartRate.ID, &heartRate.Time, &heartRate.HeartRate); err != nil {
			return HeartRateV1{}, err
	}
	return heartRate, nil
}