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

func (c *Client) GetList(params Values) (*List, error) {
	lst := new(List)
	err := c.GET("lists/show", params, &lst)
	return lst, err
}

func (c *Client) GetLists(params Values) ([]List, error) {
	var lists []List
	err := c.GET("lists/list", params, &lists)
	return lists, err
}

func (c *Client) CreateList(params Values) (*List, error) {
	lst := new(List)
	err := c.POST("lists/create", params, &lst)
	return lst, err
}

func (c *Client) DeleteList(params Values) (*List, error) {
	lst := new(List)
	err := c.POST("lists/destroy", params, &lst)
	return lst, err
}

func (c *Client) GetListMembers(params Values) (*ListMembers, error) {
	members := new(ListMembers)
	err := c.GET("lists/members", params, &members)
	return members, err
}

func (c *Client) AddListMember(params Values) (*List, error) {
	lst := new(List)
	err := c.POST("lists/members/create", params, &lst)
	return lst, err
}

func (c *Client) AddListMembers(params Values) (*List, error) {
	lst := new(List)
	err := c.POST("lists/members/create_all", params, &lst)
	return lst, err
}

func (c *Client) DeleteListMember(params Values) (*List, error) {
	lst := new(List)
	err := c.POST("lists/members/destroy", params, &lst)
	return lst, err
}

func (c *Client) DeleteListMembers(params Values) (*List, error) {
	lst := new(List)
	err := c.POST("lists/members/destroy_all", params, &lst)
	return lst, err
}
