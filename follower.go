package twtr

func (c *Client) GetFollowerIDs(params *Values) (*UserIDs, *Response, error) {
	ids := new(UserIDs)
	resp, err := c.GET("followers/ids", params, &ids)
	return ids, resp, err
}
