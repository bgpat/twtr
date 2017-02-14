package twtr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeResponse(resp *http.Response, data interface{}) error {
	if err := NewErrors(resp); err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *Client) formatURL(path string, params *Values) string {
	return fmt.Sprintf(c.URLFormat, params.ParseURL(path))
}

func (c *Client) getRequest(urlStr string, params *Values) (*http.Response, error) {
	return c.OAuthClient.Get(nil, &c.AccessToken.Credentials, params.ParseURL(urlStr), params.ToURLValues())
}

func (c *Client) postRequest(urlStr string, params *Values) (*http.Response, error) {
	return c.OAuthClient.Post(nil, &c.AccessToken.Credentials, params.ParseURL(urlStr), params.ToURLValues())
}

func (c *Client) GET(path string, params *Values, data interface{}) error {
	urlStr := c.formatURL(path, params)
	resp, err := c.getRequest(urlStr, params)
	if err != nil {
		return err
	}
	return decodeResponse(resp, data)
}

func (c *Client) POST(path string, params *Values, data interface{}) error {
	urlStr := c.formatURL(path, params)
	resp, err := c.postRequest(urlStr, params)
	if err != nil {
		return err
	}
	return decodeResponse(resp, data)
}
