package twtr_test

import (
	"testing"
)

func TestGetBlockIDs(t *testing.T) {
	ids, resp, err := client.GetBlockIDs(nil)
	if err != nil {
		t.Errorf("Failed blocks/ids: %v", err)
	} else if ids == nil {
		t.Error("Invalid user IDs")
	}
	t.Logf("UserIDs: %+v", ids)
	t.Logf("Response: %+v", resp)
}

func TestGetBlockList(t *testing.T) {
	list, resp, err := client.GetBlockList(nil)
	if err != nil {
		t.Errorf("Failed blocks/list: %v", err)
	} else if list == nil {
		t.Error("Invalid user list")
	}
	t.Logf("UserList: %+v", list)
	t.Logf("Response: %+v", resp)
}
