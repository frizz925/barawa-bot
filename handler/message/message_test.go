package message

import (
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	if !strings.Contains("apeljeruk", ProcessMessage("Apakah apel atau jeruk?")) {
		t.Error("First handler should be ChooseHandler")
	}

	if !strings.Contains("YaTidak", ProcessMessage("Apakah benar")) {
		t.Error("Second handler should be YesOrNoHandler")
	}

	if len(ProcessMessage("Tidak trigger handler")) > 0 {
		t.Error("No handler should be available")
	}
}
