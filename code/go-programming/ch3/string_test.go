package ch3

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "hello,world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	//c := s[len(s)] // panic: runtime error: index out of range [11] with length 11
	//fmt.Println(c)

	fmt.Println(s[0:5]) // hello
	fmt.Println(s[:5])  // hello
	fmt.Println(s[6:])  // world
	fmt.Println(s[:])   // hello,world

	//s[0] = 'a' // compile error: cannot assign to s[0]
}
