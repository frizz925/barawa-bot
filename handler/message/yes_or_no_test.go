package message

import (
	"testing"
)

func TestYesOrNo(t *testing.T) {

	handler := YesOrNoHandler{}
	message := "Apakah ini benar?"
	if !handler.test(message) {
		t.Error("Should pass the handler test")
	}

	message = "Test"
	if handler.test(message) {
		t.Error("Should not pass the handler test")
	}

	message = "Ini tidak akan lewat."
	if handler.test(message) {
		t.Error("Should not pass the handler test")
	}

	/*
		rand.Seed(1) // returns above 0.5
		result := handler.handle("Dummy")
		if result != "Ya" {
			t.Error("Should return 'Ya'")
		}

		rand.Seed(2) // returns below 0.5
		result = handler.handle("Dummy")
		if result != "Tidak" {
			t.Error("Should return 'Tidak'")
		}
	*/
}
