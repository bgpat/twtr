package twtr

import (
	"net/url"
)

type Values map[string]string

func (v Values) ToURLValues() url.Values {
	u := url.Values{}
	for n, s := range v {
		u[n] = []string{s}
	}
	return u
}
