package twtr

// GetCSRFToken returns auth/csrf_token
func (c *Client) GetCSRFToken(params *Params) (*Cookie, *Response, error) {
	cookie := new(Cookie)
	resp, err := c.GET("auth/csrf_token", params, cookie)
	return cookie, resp, err
}
