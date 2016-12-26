package twtr

type Images struct {
	Sizes map[string]Image `json:"sizes"`
}

type Image struct {
	H   int    `json:"h"`
	W   int    `json:"w"`
	URL string `json:"url"`
}
