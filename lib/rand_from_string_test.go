package lib

import (
	"testing"
)

func TestRandFromString(t *testing.T) {
	result := RandFromString("abc", 0)
	if result != 294 {
		panic("Result from RandFromString(\"abc\") should be 294")
	}
}
