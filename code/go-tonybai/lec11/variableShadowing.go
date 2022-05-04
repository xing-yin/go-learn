package lec11

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/4/27
 * @Desc: 演示 Go 变量遮蔽
 */

var a = 11

func foo(n int) {
	a := 1
	a += n
}

func main() {
	fmt.Println("a=", a)
	foo(5)
	fmt.Println("after calling foo, a=", a)

	var s string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
	fmt.Println(s)
}
