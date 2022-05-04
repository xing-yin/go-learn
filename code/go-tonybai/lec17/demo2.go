package lec17

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/2
 * @Desc:
 */

type T = string

func main() {
	var s string = "hi"
	var t T = s
	fmt.Printf("%T\n", t)
}
