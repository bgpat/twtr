package twtr

// GetFollowerIDs returns a cursored collection of user IDs for every user following the specified user.
// https://dev.twitter.com/rest/reference/get/followers/ids
func (c *Client) GetFollowerIDs(params *Params) (*UserIDs, *Response, error) {
	ids := new(UserIDs)
	resp, err := c.GET("followers/ids", params, ids)
	return ids, resp, err
}

// GetFollowerList returns a cursored collection of user objects for users following the specified user.
// https://dev.twitter.com/rest/reference/get/followers/list
func (c *Client) GetFollowerList(params *Params) (*UserList, *Response, error) {
	list := new(UserList)
	resp, err := c.GET("followers/list", params, &list)
	return list, resp, err
}
