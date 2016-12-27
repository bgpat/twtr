package twtr

import (
	"net/url"
	"strings"
)

type Values map[string]string

func (v *Values) ToURLValues() (u url.Values) {
	if v == nil {
		return
	}
	for n, s := range *v {
		u[n] = []string{s}
	}
	return
}

func (v *Values) ParseURL(urlStr string) string {
	if v == nil {
		return urlStr
	}
	for k, s := range *v {
		if len(k) <= 0 || k[0] != ':' {
			continue
		}
		urlStr = strings.Replace(urlStr, k, s, -1)
		delete(*v, k)
	}
	return urlStr
}
