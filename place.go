package twtr

type Place struct {
	Attributes      map[string]string `json:"attributes"`
	BoundingBox     PlaceBoundingBox  `json:"bounding_box"`
	ContainedWithin *Place            `json:"contained_within"`
	Country         string            `json:"country"`
	CountryCode     string            `json:"country_code"`
	FullName        string            `json:"full_name"`
	Geometry        PlaceBoundingBox  `json:"geometry"`
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	PlaceType       string            `json:"place_type"`
	URL             string            `json:"url"`
}

type PlaceBoundingBox struct {
	Coordinates [][4][2]float64 `json:"coordinates"`
	Type        string          `json:"type"`
}
