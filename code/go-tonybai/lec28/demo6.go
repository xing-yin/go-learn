package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc:
 */

type MyInterface1 interface {
	M1()
}

type T int

func (T) M1() {
	println("T's M1")
}

func main() {
	var t T
	var i interface{} = t
	v1, ok := i.(MyInterface1)
	if !ok {
		panic("the value of i is not MyInterface1")
	}
	v1.M1()
	fmt.Printf("th")

}
