package twtr_test

import (
	"os"
	"testing"

	"github.com/bgpat/twtr"
)

var (
	consumerKey    = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	accessToken    = os.Getenv("ACCESS_TOKEN")
	accessSecret   = os.Getenv("ACCESS_SECRET")
	client         *twtr.Client
)

func init() {
	consumer := twtr.NewCredentials(consumerKey, consumerSecret)
	token := twtr.NewCredentials(accessToken, accessSecret)
	client = twtr.NewClient(consumer, token)
}

func TestRequestTokenURL(t *testing.T) {
	url, err := client.RequestTokenURL("http://localhost/")
	if err != nil {
		t.Errorf("Failed to request authorization URL: %v", err)
	} else if url == "" {
		t.Error("Request token URL is empty")
	}
}
