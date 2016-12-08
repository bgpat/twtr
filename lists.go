package twtr

type List struct {
	Slug            string `json:"slug,omitempty"`
	Name            string `json:"name,omitempty"`
	CreatedAt       Time   `json:"created_at,omitempty"`
	URI             string `json:"uri,omitempty"`
	SubscriberCount int    `json:"subscriber_count,omitempty"`
	IDStr           string `json:"id_str,omitempty"`
	MemberCount     int    `json:"member_count,omitempty"`
	Mode            string `json:"mode,omitempty"`
	ID              int64  `json:"id,omitempty"`
	FullName        string `json:"full_name,omitempty"`
	Description     string `json:"description,omitempty"`
	User            User   `json:"user,omitempty"`
	Following       bool   `json:"following,omitempty"`
}

type ListMembers struct {
	Cursor
	Users []User `json:"users,omitempty"`
}

func (c *Client) GetLists(params Values) (*[]List, error) {
	var data []List
	err := c.GET("lists/list", params, &data)
	return &data, err
}

func (c *Client) GetList(params Values) (*List, error) {
	var data List
	err := c.GET("lists/show", params, &data)
	return &data, err
}

func (c *Client) GetListMembers(params Values) (*ListMembers, error) {
	var data ListMembers
	err := c.GET("lists/members", params, &data)
	return &data, err
}

func (c *Client) CreateList(params Values) (*List, error) {
	var data List
	err := c.POST("lists/create", params, &data)
	return &data, err
}

func (c *Client) AddListMember(params Values) (*List, error) {
	var data List
	err := c.POST("lists/members/create", params, &data)
	return &data, err
}

func (c *Client) AddListMembers(params Values) (*List, error) {
	var data List
	err := c.POST("lists/members/create_all", params, &data)
	return &data, err
}

func (c *Client) DeleteListMember(params Values) (*List, error) {
	var data List
	err := c.POST("lists/members/destroy", params, &data)
	return &data, err
}

func (c *Client) DeleteListMembers(params Values) (*List, error) {
	var data List
	err := c.POST("lists/members/destroy_all", params, &data)
	return &data, err
}
