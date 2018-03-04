package message

import "testing"

func TestChoose(t *testing.T) {
	handler := ChooseHandler{}
	message := "Ini tidak akan lewat"
	if handler.test(message) {
		t.Error("Should not pass the handler test")
	}

	message = "Test"
	if handler.test(message) {
		t.Error("Should not pass the handler test")
	}

	message = "Apakah ini akan lewat?"
	if handler.test(message) {
		t.Error("Should not pass the handler test")
	}

	message = "Apakah apel atau jeruk?"
	if !handler.test(message) {
		t.Error("Should pass the handler test")
	}

	result := handler.handle(message)
	if result != "apel" && result != "jeruk" {
		t.Error("Should return either 'apel' or 'jeruk'")
	}
}
