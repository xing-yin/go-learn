package loop_test

import (
	"fmt"
	"testing"
)

// while 循环
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

// 无限循环
func TestEndlessLoop(t *testing.T) {
	for {
		fmt.Println("EndlessLoop")
	}
}
