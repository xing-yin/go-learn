package loop__test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestWhile(t *testing.T) {
	n := 0
	for {
		t.Log(n)
		n++
	}
}
