package employee

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc: 演示-E1 和 E2 方法集合存在交集
 */

type E1 interface {
	M1()
	M2()
	M3()
}

type E2 interface {
	M1()
	M2()
	M4()
}

type T6 struct {
	E1
	E2
}

// E1 和 E2 方法集合存在交集的解决办法
func (t T6) M1() {
	println("T6 's M1")
}

func (t T6) M2() {
	println("T6 's M2")
}

func main() {
	t := T6{}
	t.M1()
	t.M2()
}
