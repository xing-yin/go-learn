package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//foo1()
	foo2()
	//foo3()
}

func foo3() (int, error) {
	// 使用 strings.Join 在数据量大时更高效
	return fmt.Println(strings.Join(os.Args[1:], " "))
}

func foo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func foo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
