package message

import (
	"crypto/rand"
	"math/big"
	"strings"
)

type YesOrNoHandler struct {
	*BaseMessageHandler
}

func (h *YesOrNoHandler) handle(message string) string {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(100)))
	seed := float64(n.Int64()) / 100
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
