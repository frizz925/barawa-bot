package lib

import (
	"testing"
)

func TestRandFromString(t *testing.T) {
	// TODO: make assertion for RandFromString without 2nd param
	result := RandFromString("abc")
	result = RandFromString("abc", 0)
	if result != 294 {
		panic("Result from RandFromString(\"abc\") should be 294")
	}
}
