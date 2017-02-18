package twtr_test

import (
	"testing"
)

func TestGetCSRFToken(t *testing.T) {
	cookie, resp, err := client.GetCSRFToken(nil)
	if err != nil {
		t.Errorf("Failed auth/csrf_token: %v", err)
	} else if cookie == nil {
		t.Error("Invalid cookie")
	}
	t.Logf("Cookie: %+v", cookie)
	t.Logf("Response: %+v", resp)
}
