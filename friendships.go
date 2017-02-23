package twtr

// Friendship represent relationship of the authenticating user.
// Values for connections can be: following, following_requested, followed_by, none, blocking, muting.
// https://dev.twitter.com/rest/reference/get/friendships/lookup
type Friendship struct {
	ID
	Name        string   `json:"name"`
	ScreenName  string   `json:"screen_name"`
	Connections []string `json:"connections"`
}

// Relationship represent relationship of between two arbitrary users.
// https://dev.twitter.com/rest/reference/get/friendships/show
type Relationship struct {
	Relationship struct {
		Target struct {
			ID
			ScreenName string `json:"screen_name"`
			Following  bool   `json:"following"`
			FollowedBy bool   `json:"followed_by"`
		} `json:"target"`
		Source struct {
			ID
			CanDM                bool   `json:"can_dm"`
			Blocking             bool   `json:"blocking"`
			Muting               bool   `json:"muting"`
			AllReplies           bool   `json:"all_replies"`
			WantRetweets         bool   `json:"want_retweets"`
			MarkedSpam           bool   `json:"marked_spam"`
			ScreenName           string `json:"screen_name"`
			Following            bool   `json:"following"`
			FollowedBy           bool   `json:"followed_by"`
			NotificationsEnabled bool   `json:"notifications_enabled"`
		} `json:"source"`
	} `json:"relationship"`
}

// GetIncomingFriendships returns a collection of numeric IDs for every user who has a pending request to follow the authenticating user.
// https://dev.twitter.com/rest/reference/get/friendships/incoming
func (c *Client) GetIncomingFriendships(params *Params) (*UserIDs, *Response, error) {
	ids := new(UserIDs)
	resp, err := c.GET("friendships/incoming", params, ids)
	return ids, resp, err
}

// GetFriendships returns the relationships of the authenticating user to the comma-separated list of up to 100 screen_names or user_ids provided.
// https://dev.twitter.com/rest/reference/get/friendships/lookup
func (c *Client) GetFriendships(params *Params) ([]Friendship, *Response, error) {
	var friendships []Friendship
	resp, err := c.GET("friendships/lookup", params, &friendships)
	return friendships, resp, err
}

// GetNoRetweetFriendshipIDs eturns a collection of user_ids that the currently authenticated user does not want to receive retweets from.
// Use POST friendships/update to set the “no retweets” status for a given user account on behalf of the current user.
// https://dev.twitter.com/rest/reference/get/friendships/no_retweets/ids
func (c *Client) GetNoRetweetFriendshipIDs(params *Params) ([]string, *Response, error) {
	var ids []string
	resp, err := c.GET("friendships/no_retweets/ids", params, &ids)
	return ids, resp, err
}

// GetOutgoingFriendships returns a collection of numeric IDs for every protected user for whom the authenticating user has a pending follow request.
// https://dev.twitter.com/rest/reference/get/friendships/outgoing
func (c *Client) GetOutgoingFriendships(params *Params) (*UserIDs, *Response, error) {
	ids := new(UserIDs)
	resp, err := c.GET("friendships/outgoing", params, &ids)
	return ids, resp, err
}

// GetFriendship returns detailed information about the relationship between two arbitrary users.
// https://dev.twitter.com/rest/reference/get/friendships/show
func (c *Client) GetFriendship(params *Params) (*Relationship, *Response, error) {
	relation := new(Relationship)
	resp, err := c.GET("friendships/show", params, relation)
	return relation, resp, err
}

// Follow allows the authenticating users to follow the user specified in the ID parameter.
// https://dev.twitter.com/rest/reference/post/friendships/create
func (c *Client) Follow(params *Params) (*User, *Response, error) {
	user := new(User)
	resp, err := c.POST("friendships/create", params, user)
	return user, resp, err
}

// Unfollow allows the authenticating user to unfollow the user specified in the ID parameter.
// https://dev.twitter.com/rest/reference/post/friendships/destroy
func (c *Client) Unfollow(params *Params) (*User, *Response, error) {
	user := new(User)
	resp, err := c.POST("friendships/destroy", params, user)
	return user, resp, err
}

// UpdateFriendship allows one to enable or disable retweets and device notifications from the specified user.
// https://dev.twitter.com/rest/reference/post/friendships/update
func (c *Client) UpdateFriendship(params *Params) (*Relationship, *Response, error) {
	relation := new(Relationship)
	resp, err := c.POST("friendships/update", params, relation)
	return relation, resp, err
}
