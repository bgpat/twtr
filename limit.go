package twtr

type Limit struct {
	Limit     int  `json:"limit"`
	Remaining int  `json:"remaining"`
	Reset     Time `json:"reset"`
}
