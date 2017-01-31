package twtr

type TrendLocation struct {
	Name        string         `json:"name"`
	CountryCode string         `json:"country_code"`
	URL         string         `json:"url"`
	WOEID       int            `json:"woeid"`
	PlaceType   TrendPlaceType `json:"place_type"`
	ParentID    int            `json:"parent_id"`
	Country     string         `json:"country"`
}

type TrendPlaceType struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}
