package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/8
 * @Desc:
 */

type Interface interface {
	M1()
	M2()
}

type Interface2 interface {
	M1()
	M2()
}

type T2 struct{}

func (t T2) M1() {
}

func (t *T2) M2() {
}

func main() {
	//var t T2
	//var pt *T2
	//var i Interface
	//
	//i = pt
	//i = t // Type does not implement 'Interface' as the 'M2' method has a pointer receiver
}
