package message

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Frizz925/barawa-bot/lib"
)

type ChooseHandler struct{}

func (h *ChooseHandler) handle(message string, params ...interface{}) string {
	prefix := "apakah "
	delimiters := []string{" atau ", ","}
	prefixLength := len(prefix)

	message = message[prefixLength:]
	choices := SplitToChoices(message, delimiters)
	if len(choices) <= 0 {
		return ""
	}
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

func SplitToChoices(text string, delimiters []string) []string {
	regexpPattern := fmt.Sprintf("(%s)", strings.Join(delimiters, "|"))
	matcher := regexp.MustCompile(regexpPattern)
	if !matcher.MatchString(text) {
		return nil
	}
	matches := matcher.Split(text, -1)
	return matches
}
