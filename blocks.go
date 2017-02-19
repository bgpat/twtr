package twtr

// GetBlockIDs returns an array of numeric user ids the authenticating user is blocking.
// https://dev.twitter.com/rest/reference/get/blocks/ids
func (c *Client) GetBlockIDs(params *Params) (*UserIDs, *Response, error) {
	ids := new(UserIDs)
	resp, err := c.GET("blocks/ids", params, ids)
	return ids, resp, err
}

// GetBlockList returns a collection of user objects that the authenticating user is blocking.
// https://dev.twitter.com/rest/reference/get/blocks/list
func (c *Client) GetBlockList(params *Params) (*UserList, *Response, error) {
	list := new(UserList)
	resp, err := c.GET("blocks/list", params, list)
	return list, resp, err
}

// Block blocks the specified user from following the authenticating user.
// In addition the blocked user will not show in the authenticating users mentions or timeline (unless retweeted by another user).
// If a follow or friend relationship exists it is destroyed.
// https://dev.twitter.com/rest/reference/post/blocks/create
func (c *Client) Block(params *Params) (*User, *Response, error) {
	user := new(User)
	resp, err := c.POST("blocks/create", params, user)
	return user, resp, err
}

// Unblock un-blocks the user specified in the ID parameter for the authenticating user.
// Returns the un-blocked user in the requested format when successful.
// If relationships existed before the block was instated, they will not be restored.
// https://dev.twitter.com/rest/reference/post/blocks/destroy
func (c *Client) Unblock(params *Params) (*User, *Response, error) {
	user := new(User)
	resp, err := c.POST("blocks/destroy", params, user)
	return user, resp, err
}
