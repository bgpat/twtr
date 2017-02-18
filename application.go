package twtr

type RateLimitStatus struct {
	RateLimitContext RateLimitContext            `json:"rate_limit_context"`
	Resources        map[string]map[string]Limit `json:"resources"`
}

type RateLimitContext struct {
	AccessToken string `json:"access_token"`
	Application string `json:"application"`
}

func (c *Client) GetRateLimitStatus(params *Params) (*RateLimitStatus, *Response, error) {
	status := new(RateLimitStatus)
	resp, err := c.GET("application/rate_limit_status", params, status)
	return status, resp, err
}
