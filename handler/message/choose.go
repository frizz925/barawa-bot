package message

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Frizz925/barawa-bot/lib"
)

// ChooseHandler Placeholder comment
type ChooseHandler struct{}

func (h *ChooseHandler) handle(message string, params ...interface{}) string {
	prefix := "apakah "
	delimiters := []string{" atau ", ","}
	prefixLength := len(prefix)

	message = message[prefixLength:]
	choices := splitToChoices(message, delimiters)
	choicesLength := len(choices)

	idx := lib.RandFromString(message, params...) % int64(choicesLength)
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

func splitToChoices(text string, delimiters []string) []string {
	regexpPattern := fmt.Sprintf("(%s)", strings.Join(delimiters, "|"))
	matcher := regexp.MustCompile(regexpPattern)
	matches := matcher.Split(text, -1)
	return matches
}
