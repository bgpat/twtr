package twtr

// GetFavorites returns the 20 most recent Tweets favorited by the authenticating or specified user.
// https://dev.twitter.com/rest/reference/get/favorites/list
func (c *Client) GetFavorites(params *Params) ([]Tweet, *Response, error) {
	var tweets []Tweet
	resp, err := c.GET("favorites/list", params, &tweets)
	return tweets, resp, err
}

// Favorite favorites the status specified in the ID parameter as the authenticating user.
// Returns the favorite status when successful.
// https://dev.twitter.com/rest/reference/post/favorites/create
func (c *Client) Favorite(params *Params) (*Tweet, *Response, error) {
	tweet := new(Tweet)
	resp, err := c.POST("favorites/create", params, tweet)
	return tweet, resp, err
}

// Unfavorite un-favorites the status specified in the ID parameter as the authenticating user.
// Returns the un-favorited status in the requested format when successful.
// https://dev.twitter.com/rest/reference/post/favorites/destroy
func (c *Client) Unfavorite(params *Params) (*Tweet, *Response, error) {
	tweet := new(Tweet)
	resp, err := c.POST("favorites/destroy", params, tweet)
	return tweet, resp, err
}
