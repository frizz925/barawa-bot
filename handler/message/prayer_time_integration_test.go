package message

import (
	"strings"
	"testing"
)

func TestIntegrationPrayerTime(t *testing.T) {
	h := PrayerTimeHandler{}
	response := h.handle("")
	if strings.Contains(response, ":(") {
		t.Error("Integration test failed")
	}
}
