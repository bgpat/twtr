package twtr

import (
	"net/http"
)

func (c *Client) NewUserStream(params *Params) (*Streaming, error) {
	return c.NewStreaming(http.MethodGet, c.UserStreamURL, params)
}
