package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//func1()

	//arrayType()
}

func arrayType() {
	q := [3]int{}
	//q = [4]int{} // Cannot use '[4]int{}' (type [4]int) as the type [3]int
	fmt.Println(q)
}

func func1() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	w := [...]int{1, 2}
	fmt.Println(q[2])
	fmt.Println(r[2])
	fmt.Println(w[1])
}
