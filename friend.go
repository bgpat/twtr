package twtr

func (c *Client) GetFriendsIDs(params *Values) (*UserIDs, *Response, error) {
	ids := new(UserIDs)
	resp, err := c.GET("friends/ids", params, &ids)
	return ids, resp, err
}
