package utils_test

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/obochurkin/go-scylladb/utils"
)

func TestIsValidUUID(t *testing.T) {
	expected, _ := gocql.ParseUUID("a622803a-1ce3-4840-ae49-fdd6bbcf30e1")

	tests := []struct {
		name     string
		input    string
		expected gocql.UUID
		wantErr  bool
	}{
		{
			name:     "Valid UUID",
			input:    "a622803a-1ce3-4840-ae49-fdd6bbcf30e1",
			expected: expected,
			wantErr:  false,
		},
		{
			name:     "Invalid UUID",
			input:    "invalid-uuid",
			expected: gocql.UUID{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.IsValidUUID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.expected {
				t.Errorf("IsValidUUID() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
