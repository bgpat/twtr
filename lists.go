package twitter

import (
	"net/url"
)

type List struct {
	Slug            string `json:"slug,omitempty"`
	Name            string `json:"name,omitempty"`
	CreatedAt       Time   `json:"created_at,omitempty"`
	URI             string `json:"uri,omitempty"`
	SubscriberCount int    `json:"subscriber_count,omitempty"`
	IDStr           string `json:"id_str,omitempty"`
	MemberCount     int    `json:"member_count,omitempty"`
	Mode            string `json:"mode,omitempty"`
	ID              uint64 `json:"id,omitempty"`
	FullName        string `json:"full_name,omitempty"`
	Description     string `json:"description,omitempty"`
	User            User   `json:"user,omitempty"`
	Following       bool   `json:"following,omitempty"`
}

type ListMembers struct {
	Cursor
	Users []User `json:"users,omitempty"`
}

func (c *Client) GetLists(params url.Values) (*[]List, error) {
	var data []List
	err := c.GET("lists/list", params, &data)
	return &data, err
}

func (c *Client) GetList(params url.Values) (*List, error) {
	var data List
	err := c.GET("lists/show", params, &data)
	return &data, err
}

func (c *Client) GetListMembers(params url.Values) (*ListMembers, error) {
	var data ListMembers
	err := c.GET("lists/members", params, &data)
	return &data, err
}
