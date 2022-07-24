package ch5

import (
	"fmt"
	"golang.org/x/net/html"
	"testing"
)

func square(n int) int {
	return n * n
}

func TestName(t *testing.T) {
	//var f func(int) int
	//f(3) // 此处f的值为nil, 会引起panic错误

	f1 := square
	fmt.Println(f1(3))
}

// 函数是可以作为参数的
func forEachNode(n *html.Node, per, post func(n *html.Node)) {

}
