package message

import (
	"math/rand"
	"strings"
)

type ChooseHandler struct{}

func (h *ChooseHandler) handle(message string) string {
	prefix := "apakah "
	delimiter := " atau "

	prefixLength := len(prefix)

	message = message[prefixLength:]
	choices := strings.Split(message, delimiter)
	choicesLength := len(choices)

	choice := choices[rand.Intn(choicesLength)]
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
