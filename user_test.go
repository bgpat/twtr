package twtr_test

import (
	"testing"

	"github.com/bgpat/twtr"
)

const (
	userID = "783214"
)

func TestGetUser(t *testing.T) {
	user, resp, err := client.GetUser(&twtr.Params{"user_id": userID})
	if err != nil {
		t.Errorf("Failed account/verify_credentials: %v", err)
	} else if user == nil || user.IDStr != userID {
		t.Error("Invalid user")
	}
	t.Logf("User: %+v", user)
	t.Logf("Response: %+v", resp)
}
