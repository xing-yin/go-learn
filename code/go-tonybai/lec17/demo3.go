package lec17

import (
	"fmt"
	"unsafe"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/2
 * @Desc:
 */

type Book struct {
	Title  string
	Pages  int
	Indexs map[string]int
}

type Empty struct{}

func main() {
	var b Book
	b.Title = "Go Programming"

	var s Empty
	fmt.Println(unsafe.Sizeof(s))

	//var c = make(chan Empty)
	//c <- Empty{}
}
