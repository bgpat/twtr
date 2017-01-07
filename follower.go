package twtr

func (c *Client) GetFollowerIDs(params *Values) (*UserIDs, error) {
	ids := new(UserIDs)
	err := c.GET("followers/ids", params, &ids)
	return ids, err
}
