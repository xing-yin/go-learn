package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/8
 * @Desc:
 */

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

//// 等价于
//func Get(t T) int{
//	return t.a
//}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

//// 等价于
//func Set(t *T, a int) int {
//	t.a = a
//	return t.a
//}

func main() {
	var t T
	t.Get()
	(&t).Set(1)

	var t2 T
	f1 := (*T).Set
	f2 := T.Get
	fmt.Printf("the type of f1 is %T\n", f1)
	fmt.Printf("the type of f2 is %T\n", f2)
	f1(&t2, 3)
	fmt.Println(f2(t2))
}