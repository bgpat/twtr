package twtr_test

import (
	"testing"
)

func TestGetRateLimitStatus(t *testing.T) {
	status, resp, err := client.GetRateLimitStatus(nil)
	if err != nil {
		t.Errorf("Failed application/rate_limit_status: %v", err)
	} else if status == nil {
		t.Error("Invalid status")
	}
	t.Logf("Status: %+v", status)
	t.Logf("Response: %+v", resp)
}
