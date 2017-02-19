package twtr_test

import (
	"testing"
)

func TestGetFavorites(t *testing.T) {
	tweets, resp, err := client.GetFavorites(nil)
	if err != nil {
		t.Errorf("Failed favorites/list: %v", err)
	} else if tweets == nil {
		t.Error("Invalid tweets")
	}
	t.Logf("Tweets: %+v", tweets)
	t.Logf("Response: %+v", resp)
}
