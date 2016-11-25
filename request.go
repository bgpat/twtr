package twtr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Decode(resp *http.Response, data interface{}) error {
	if err := NewErrors(resp); err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *Client) ParseURL(path string, params url.Values) string {
	for k, v := range params {
		if len(k) <= 0 || len(v) <= 0 || k[0] != ':' {
			continue
		}
		path = strings.Replace(path, k, strings.Join(v, ","), -1)
		delete(params, k)
	}
	return fmt.Sprintf(c.URLFormat, path)
}

func (c *Client) GET(path string, params url.Values, data interface{}) error {
	uri := c.ParseURL(path, params)
	resp, err := c.OAuthClient.Get(nil, c.AccessToken, uri, params)
	if err != nil {
		return err
	}
	return Decode(resp, data)
}
