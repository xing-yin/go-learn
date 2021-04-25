package main

import "fmt"

/**
函数-多返回值
Go 内建多返回值支持。这个特性在 Go 语言中经常被用到，例如用来同时返回一个函数的结果和错误信息
*/

//(int, int) 在这个函数中标志着这个函数返回 2 个 int
func vals() (int, string) {
	return 1, "a"
}

func main() {

	//这里我们通过多赋值操作来使用这两个不同的返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//如果你仅仅想返回值的一部分的话，你可以使用空白定义符 _。
	_, c := vals()
	fmt.Println(c)

}
