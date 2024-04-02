package buildings

import (
	"github.com/gocql/gocql"
)

const (
	InsertBuilding       = `INSERT INTO buildings (name, city, built_year, height_meters) VALUES (?, ?, ?, ?);`
	SelectBuildingByCity = `SELECT name, city, built_year, height_meters FROM mykeyspace.buildings_by_city WHERE city = ?;`
	SelectAllBuildings   = `SELECT name, city, built_year, height_meters FROM buildings;`
)

func InsertBuildingQuery(session *gocql.Session, data *AddBuildingDto) error {
	return session.Query(InsertBuilding, data.Name, data.City, data.BuiltYear, data.HeightMeters).Exec()
}

func GetBuildingByCity(session *gocql.Session, data *GetBuildingQueryParam) ([]Building, error) {
	var buildings []Building
	var query string
	var queryParams []interface{}

	if data.City != nil {
		query = SelectBuildingByCity
		queryParams = append(queryParams, *data.City)
	} else {
		query = SelectAllBuildings
	}

	iter := session.Query(query, queryParams...).Iter()

	for {
		var building Building

		if !iter.Scan(&building.Name, &building.City, &building.BuiltYear, &building.HeightMeters) {
			break
		}
		buildings = append(buildings, building)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return buildings, nil
}
