package message

import (
	"fmt"
	"testing"
)

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

	MessageTest(t, &handler, "Apakah apel atau jeruk?", []string{"apel", "jeruk"})
	MessageTest(t, &handler, "Apakah apel, jeruk, atau mangga?", []string{"apel", "jeruk", "mangga"})
	MessageTest(t, &handler, "apakah a, b, atau c", []string{"a", "b", "c"})
}

func MessageTest(t *testing.T, handler MessageHandler, message string, choices []string) {
	if !handler.test(message) {
		t.Error(fmt.Sprintf("\"%s\" should pass the handler test", message))
	}

	result := handler.handle(message, 0)
	for _, choice := range choices {
		if result == choice {
			return
		}
	}
	t.Error(fmt.Sprintf("No matching choice returned from \"%s\"", message))
}
