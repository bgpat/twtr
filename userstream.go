package twtr

import (
	"net/http"
)

func (c *Client) NewUserStream(params *Values) (*Streaming, error) {
	return c.NewStreaming(http.MethodGet, c.UserStreamURL, params)
}
