package twtr_test

import (
	"testing"
)

func TestGetFollowerIDs(t *testing.T) {
	ids, resp, err := client.GetFollowerIDs(nil)
	if err != nil {
		t.Errorf("Failed followers/ids: %v", err)
	} else if ids == nil {
		t.Error("Invalid userIDs")
	}
	t.Logf("UserIDs: %+v", ids)
	t.Logf("Response: %+v", resp)
}

func TestGetFollowerList(t *testing.T) {
	list, resp, err := client.GetFollowerList(nil)
	if err != nil {
		t.Errorf("Failed followers/list: %v", err)
	} else if list == nil {
		t.Error("Invalid user list")
	}
	t.Logf("UserList: %+v", list)
	t.Logf("Response: %+v", resp)
}
