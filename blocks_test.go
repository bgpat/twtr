package twtr_test

import (
	"testing"
)

func TestGetBlockUserIDs(t *testing.T) {
	ids, resp, err := client.GetBlockUserIDs(nil)
	if err != nil {
		t.Errorf("Failed blocks/ids: %v", err)
	} else if ids == nil {
		t.Error("Invalid user IDs")
	}
	t.Logf("UserIDs: %+v", ids)
	t.Logf("Response: %+v", resp)
}
