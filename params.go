package twtr

import (
	"net/url"
	"strings"
)

type Params map[string]string

func (p *Params) ToURLValues() (u url.Values) {
	if p == nil {
		return
	}
	u = url.Values{}
	for n, s := range *p {
		u[n] = []string{s}
	}
	return
}

func (p *Params) ParseURL(urlStr string) string {
	if p == nil {
		return urlStr
	}
	for k, s := range *p {
		if len(k) <= 0 || k[0] != ':' {
			continue
		}
		urlStr = strings.Replace(urlStr, k, s, -1)
		delete(*p, k)
	}
	return urlStr
}

func (p *Params) AddNextCursor(cursor Cursor) *Params {
	(*p)["cursor"] = cursor.NextCursorStr
	return p
}

func (p *Params) AddPreviousCursor(cursor Cursor) *Params {
	(*p)["cursor"] = cursor.PreviousCursorStr
	return p
}
