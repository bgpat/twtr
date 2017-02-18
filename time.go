package twtr

import (
	"strconv"
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

func (t *Time) ParseUnixTime(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(i, 0)
	return nil
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	for _, c := range b {
		if c < '0' || '9' < c {
			t.Time, err = time.Parse(layout, string(b))
			return
		}
	}
	err = t.ParseUnixTime(string(b))
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(layout)), nil
}
