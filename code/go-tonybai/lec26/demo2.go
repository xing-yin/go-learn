package employee

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc:
 */

type interface1 interface {
	M1()
}

type interface2 interface {
	M1()
	M2()
}

type interface3 interface {
	interface1
	interface2
}

type interface4 interface {
	interface2
	M2()
}

func main() {

}
