package twtr

func (c *Client) GetFriendsIDs(params *Values) (*UserIDs, error) {
	ids := new(UserIDs)
	err := c.GET("friends/ids", params, &ids)
	return ids, err
}
