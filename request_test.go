package twtr_test

import (
	"testing"
)

func TestGET(t *testing.T) {
	var user interface{}
	resp, err := client.GET("account/verify_credentials", nil, &user)
	if err != nil {
		t.Errorf("Failed account/verify_credentials: %v", err)
	}
	t.Logf("User: %+v\n", user)
	t.Logf("Response: %+v\n", resp)
}
