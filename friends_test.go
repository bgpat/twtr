package twtr_test

import (
	"testing"

	"github.com/bgpat/twtr"
)

func TestGetFriendIDs(t *testing.T) {
	ids, resp, err := client.GetFriendIDs(nil)
	if err != nil {
		t.Errorf("Failed friends/ids: %v", err)
	} else if ids == nil {
		t.Error("Invalid userIDs")
	}
	t.Logf("UserIDs: %+v", ids)
	t.Logf("Response: %+v", resp)
}

func TestGetFriendList(t *testing.T) {
	list, resp, err := client.GetFriendList(nil)
	if err != nil {
		t.Errorf("Failed friends/list: %v", err)
	} else if list == nil {
		t.Error("Invalid user list")
	}
	t.Logf("UserList: %+v", list)
	t.Logf("Response: %+v", resp)
}

func TestGetFriendFollowingIDs(t *testing.T) {
	params := &twtr.Params{"screen_name": "Twitter"}
	ids, resp, err := client.GetFriendFollowingIDs(params)
	if err != nil {
		t.Errorf("Failed friends/ids: %v", err)
	} else if ids == nil {
		t.Error("Invalid userIDs")
	}
	t.Logf("UserIDs: %+v", ids)
	t.Logf("Response: %+v", resp)
}

func TestGetFriendFollowingList(t *testing.T) {
	params := &twtr.Params{"screen_name": "Twitter"}
	list, resp, err := client.GetFriendFollowingList(params)
	if err != nil {
		t.Errorf("Failed friends/list: %v", err)
	} else if list == nil {
		t.Error("Invalid user list")
	}
	t.Logf("UserList: %+v", list)
	t.Logf("Response: %+v", resp)
}
