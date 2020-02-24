package gas

type Station struct {
	ID     string
	Name   string
	Brand  string
	Street string
	Place  string
	Diesel float64
	E5     float64
	E10    float64
}

func NewStationFromJson(station map[string]interface{}) Station {
	return Station{
		ID:     station["id"].(string),
		Name:   station["name"].(string),
		Brand:  station["brand"].(string),
		Street: station["street"].(string),
		Place:  station["place"].(string),
		Diesel: station["diesel"].(float64),
		E5:     station["e5"].(float64),
		E10:    station["e10"].(float64),
	}
}
