package twtr

// GetFriendIDs returns a cursored collection of user IDs for every user the specified user is following (otherwise known as their “friends”).
// https://dev.twitter.com/rest/reference/get/friends/ids
func (c *Client) GetFriendIDs(params *Params) (*UserIDs, *Response, error) {
	ids := new(UserIDs)
	resp, err := c.GET("friends/ids", params, &ids)
	return ids, resp, err
}
