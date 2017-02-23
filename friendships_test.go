package twtr_test

import (
	"testing"

	"github.com/bgpat/twtr"
)

func TestGetIncomingFriendships(t *testing.T) {
	friendships, resp, err := client.GetIncomingFriendships(nil)
	if err != nil {
		t.Errorf("Failed friendships/outgoing: %v", err)
	} else if friendships == nil {
		t.Error("Invalid friendships")
	}
	t.Logf("Friendships: %+v", friendships)
	t.Logf("Response: %+v", resp)
}

func TestGetFriendships(t *testing.T) {
	params := &twtr.Params{"screen_name": "Twitter"}
	friendships, resp, err := client.GetFriendships(params)
	if err != nil {
		t.Errorf("Failed friendships/lookup: %v", err)
	} else if friendships == nil {
		t.Error("Invalid friendships")
	}
	t.Logf("Friendships: %+v", friendships)
	t.Logf("Response: %+v", resp)
}

func TestGetNoRetweetFriendshipIDs(t *testing.T) {
	ids, resp, err := client.GetNoRetweetFriendshipIDs(nil)
	if err != nil {
		t.Errorf("Failed friendships/no_retweets/ids: %v", err)
	} else if ids == nil {
		t.Error("Invalid userIDs")
	}
	t.Logf("UserIDs: %+v", ids)
	t.Logf("Response: %+v", resp)
}

func TestGetOutgoingFriendships(t *testing.T) {
	friendships, resp, err := client.GetOutgoingFriendships(nil)
	if err != nil {
		t.Errorf("Failed friendships/outgoing: %v", err)
	} else if friendships == nil {
		t.Error("Invalid friendships")
	}
	t.Logf("Friendships: %+v", friendships)
	t.Logf("Response: %+v", resp)
}

func TestGetFriendship(t *testing.T) {
	params := &twtr.Params{
		"source_screen_name": "Twitter",
		"target_screen_name": "twitterapi",
	}
	friendship, resp, err := client.GetFriendship(params)
	if err != nil {
		t.Errorf("Failed friendships/show %v", err)
	} else if friendship == nil {
		t.Error("Invalid friendship")
	}
	t.Logf("Friendship: %+v", friendship)
	t.Logf("Response: %+v", resp)
}
