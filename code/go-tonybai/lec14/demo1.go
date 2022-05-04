package lec13

import "fmt"

/**
* @Author: Alan Yin
* @Date: 2022/4/30
* @Desc:
 */

type myInt int

const n myInt = 12

// error
//const m int =  n + 1
const m int = int(n) + 1

const k = 2

const j = 1333333333

func main() {

	// error
	//a := 5
	// fmt.Println(a + n)

	// ok
	var b myInt = 5
	fmt.Println(b + k)

	//var c int8 = 2
	//d := c + j // 编译器报错：constant 1333333333 overflows int8
}
