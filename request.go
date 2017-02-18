package twtr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeResponse(resp *http.Response, data interface{}) (*Response, error) {
	if err := NewErrors(resp); err != nil {
		return nil, err
	}
	r := new(Response)
	err := r.ParseResponse(resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(data)
	return r, err
}

func (c *Client) formatURL(path string, params *Params) string {
	return fmt.Sprintf(c.URLFormat, params.ParseURL(path))
}

func (c *Client) getRequest(urlStr string, params *Params) (*http.Response, error) {
	return c.OAuthClient.Get(nil, &c.AccessToken.Credentials, params.ParseURL(urlStr), params.ToURLValues())
}

func (c *Client) postRequest(urlStr string, params *Params) (*http.Response, error) {
	return c.OAuthClient.Post(nil, &c.AccessToken.Credentials, params.ParseURL(urlStr), params.ToURLValues())
}

func (c *Client) GET(path string, params *Params, data interface{}) (*Response, error) {
	urlStr := c.formatURL(path, params)
	resp, err := c.getRequest(urlStr, params)
	if err != nil {
		return nil, err
	}
	return decodeResponse(resp, data)
}

func (c *Client) POST(path string, params *Params, data interface{}) (*Response, error) {
	urlStr := c.formatURL(path, params)
	resp, err := c.postRequest(urlStr, params)
	if err != nil {
		return nil, err
	}
	return decodeResponse(resp, data)
}
