package message

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestPrayerTime(t *testing.T) {
	h := PrayerTimeHandler{}
	if !h.test("Kapan waktu azan") {
		t.Error("Should pass the test")
	}
	if !h.test("Kapan waktu buka puasa") {
		t.Error("Should pass the test")
	}
	if h.test("Random message") {
		t.Error("Should not pass the test")
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadFile(filepath.Join(dir, "../../fixtures/prayer/Schedule201804.json"))
	if err != nil {
		t.Error(err)
	}
	x := Response{}
	json.Unmarshal(data, &x)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	res := h.handleResponse(&x, time.Date(2018, 05, 18, 0, 0, 0, 0, loc))
	if !strings.Contains(res, "Maghrib: 17:47") {
		t.Error("Invalid response")
	}
}
