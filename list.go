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

func (c *Client) GetList(params *Params) (*List, *Response, error) {
	list := new(List)
	resp, err := c.GET("lists/show", params, &list)
	return list, resp, err
}

func (c *Client) GetLists(params *Params) ([]List, *Response, error) {
	var lists []List
	resp, err := c.GET("lists/list", params, &lists)
	return lists, resp, err
}

func (c *Client) CreateList(params *Params) (*List, *Response, error) {
	list := new(List)
	resp, err := c.POST("lists/create", params, &list)
	return list, resp, err
}

func (c *Client) DeleteList(params *Params) (*List, *Response, error) {
	list := new(List)
	resp, err := c.POST("lists/destroy", params, &list)
	return list, resp, err
}

func (c *Client) GetListMembers(params *Params) (*ListMembers, *Response, error) {
	members := new(ListMembers)
	resp, err := c.GET("lists/members", params, &members)
	return members, resp, err
}

func (c *Client) AddListMember(params *Params) (*List, *Response, error) {
	list := new(List)
	resp, err := c.POST("lists/members/create", params, &list)
	return list, resp, err
}

func (c *Client) AddListMembers(params *Params) (*List, *Response, error) {
	list := new(List)
	resp, err := c.POST("lists/members/create_all", params, &list)
	return list, resp, err
}

func (c *Client) DeleteListMember(params *Params) (*List, *Response, error) {
	list := new(List)
	resp, err := c.POST("lists/members/destroy", params, &list)
	return list, resp, err
}

func (c *Client) DeleteListMembers(params *Params) (*List, *Response, error) {
	list := new(List)
	resp, err := c.POST("lists/members/destroy_all", params, &list)
	return list, resp, err
}
