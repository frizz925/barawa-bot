package message

import (
	"strings"

	"github.com/Frizz925/barawa-bot/lib"
)

type YesOrNoHandler struct {
	*BaseMessageHandler
}

func (h *YesOrNoHandler) handle(message string) string {
	seed := float64(lib.RandFromString(message)%100) / 100
	if seed > 0.5 {
		return "Ya"
	}
	return "Tidak"
}

func (h *YesOrNoHandler) test(message string) bool {
	prefix := "apakah "
	prefixLength := len(prefix)
	if len(message) < prefixLength {
		return false
	}
	messagePrefix := strings.ToLower(message[:prefixLength])
	return messagePrefix == prefix
}
