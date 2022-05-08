package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/8
 * @Desc: 选择不同的 receiver 类型对原类型实例的影响
 */

type T struct {
	a int
}

func (t T) M11() {
	t.a = 10
}

func (t *T) M22() {
	t.a = 11
}

func main() {
	var t T
	println(t.a) // 0

	t.M11()
	println(t.a) // 0 ==> 方法 M11 由于使用了 T2 作为 receiver 参数类型，它在方法体中修改的仅仅是 T2 类型实例 t 的副本，原实例并没有受到影响。因此 M11 调用后，输出 t.a 的值仍为 0。

	p := &t
	p.M22()
	println(t.a) // 11 ==> 由于使用了 *T2 作为 receiver 参数类型，它在方法体中通过 t 修改的是实例本身，因此 M22 调用后，t.a 的值变为了 11。
}
