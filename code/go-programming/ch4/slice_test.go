package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var s []int
	fmt.Printf("len:[%d],cap:[%d]\n", len(s), cap(s))
	s = nil
	fmt.Printf("len:[%d],cap:[%d]\n", len(s), cap(s))
	s = []int(nil)
	fmt.Printf("len:[%d],cap:[%d]\n", len(s), cap(s))
	s = []int{}
	fmt.Printf("len:[%d],cap:[%d]\n", len(s), cap(s))
}

func equalSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
