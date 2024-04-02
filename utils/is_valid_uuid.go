package utils

import (
	"github.com/gocql/gocql"
)

func IsValidUUID(input string) (gocql.UUID, error) {
	uuid, err := gocql.ParseUUID(input)
	if err != nil {
		return gocql.UUID{}, err // Return a zero UUID and the error
	}

	return uuid, nil // Return the parsed UUID and no error
}