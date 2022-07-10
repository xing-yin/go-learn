package ch3

import "testing"

func TestBoolToInt(t *testing.T) {
	boolToInt(true)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
