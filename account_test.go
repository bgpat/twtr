package twtr_test

import (
	"testing"
)

func TestVerifyCredentials(t *testing.T) {
	user, err := client.VerifyCredentials(nil)
	if err != nil {
		t.Errorf("Failed account/verify_credentials: %v", err)
	} else if user == nil {
		t.Error("Invalid user")
	}
}
