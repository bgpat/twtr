package twtr

import (
	"time"
)

type Time struct {
	time.Time
}

type TimeZone struct {
	Name       string `json:"name"`
	UTCOffset  int    `json:"utc_offset"`
	TZInfoName string `json:"tzinfo_name"`
}

var layout = `"` + time.RubyDate + `"`

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse(layout, string(b))
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(layout)), nil
}
