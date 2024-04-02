package buildings

type Building struct {
	Name         string     `json:"name"`
	City         string     `json:"city"`
	BuiltYear    int        `json:"built_year"`
	HeightMeters int        `json:"height_meters"`
}
