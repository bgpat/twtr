package twitter

import (
	"time"
)

type Time struct {
	time.Time
}

var layout = `"` + time.RubyDate + `"`

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse(layout, string(b))
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(layout)), nil
}
