package twtr

import (
	"net/url"
	"strings"
)

type Values map[string]string

func (v Values) ToURLValues() url.Values {
	u := url.Values{}
	for n, s := range v {
		u[n] = []string{s}
	}
	return u
}

func (v Values) ParseURL(urlStr string) string {
	for k, s := range v {
		if len(k) <= 0 || k[0] != ':' {
			continue
		}
		urlStr = strings.Replace(urlStr, k, s, -1)
		delete(v, k)
	}
	return urlStr
}
