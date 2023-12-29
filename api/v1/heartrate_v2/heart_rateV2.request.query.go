package heartrate_v2

import "time"

// ISO 8601 (e.g., 2006-01-02T15:04:05Z).
type HeartRateV2QueryParams struct {
	StartDate time.Time `form:"start-date" validate:"required,datetime,max=50" time_format:"2006-01-02T15:04:05Z"`
	EndDate   time.Time `form:"end-date,omitempty" validate:"omitempty,datetime,max=50" time_format:"2006-01-02T15:04:05Z"`
}
