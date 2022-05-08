package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/8
 * @Desc:
 */

type T1 struct {
	a int
}

func (t T1) M11() {
	t.a = 10
}

func (t *T1) M22() {
	t.a = 11
}

func main() {
	var t1 T1
	println(t1.a) // 0
	t1.M11()
	println(t1.a) // 0
	t1.M22()      // ==> Go 编译器自动将 t1.M22() 转换为 (*t1).M22()
	println(t1.a) // 11

	var t2 = &T1{}
	println(t2.a) // 0
	t2.M11()      // ==> Go 编译器自动将 t2.M11() 转换为 (*t2).M11()
	println(t2.a) // 0
	t2.M22()
	println(t2.a) // 11
}
