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
		result := handler.handle("Apel")
		if result != "Ya" {
			t.Error("Should return 'Ya'")
		}

		result = handler.handle("Dummy")
		if result != "Tidak" {
			t.Error("Should return 'Tidak'")
		}
	*/
}
