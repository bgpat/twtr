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
