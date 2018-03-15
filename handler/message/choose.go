package message

import (
	"strings"

	"github.com/Frizz925/barawa-bot/lib"
)

type ChooseHandler struct{}

func (h *ChooseHandler) handle(message string) string {
	prefix := "apakah "
	delimiter := " atau "

	prefixLength := len(prefix)

	message = message[prefixLength:]
	choices := strings.Split(message, delimiter)
	choicesLength := len(choices)

	idx := lib.RandFromString(message) % int64(choicesLength)
	choice := choices[idx]
	choice = strings.TrimSpace(choice)
	choice = strings.Trim(choice, "?!.")

	return choice
}

func (h *ChooseHandler) test(message string) bool {
	prefix := "apakah "
	prefixLength := len(prefix)
	if len(message) < prefixLength {
		return false
	}
	messagePrefix := strings.ToLower(message[:prefixLength])
	if messagePrefix != prefix {
		return false
	}
	delimiter := " atau "
	return strings.Contains(message, delimiter)
}
