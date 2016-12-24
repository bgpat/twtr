package twtr

type List struct {
	ID
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	CreatedAt       Time   `json:"created_at"`
	URI             string `json:"uri"`
	SubscriberCount int    `json:"subscriber_count"`
	MemberCount     int    `json:"member_count"`
	Mode            string `json:"mode"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	User            User   `json:"user"`
	Following       bool   `json:"following"`
}

type ListMembers struct {
	Cursor
	Users []User `json:"users"`
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

func (c *Client) DeleteList(params Values) (*List, error) {
	var data List
	err := c.POST("lists/destroy", params, &data)
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
