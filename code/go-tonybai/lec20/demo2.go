package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

func case4() int {
	println("eval case4 expr")
	return 1
}

func case5() int {
	println("eval case5 expr")
	return 2
}

func switchexpr2() int {
	println("eval switch expr")
	return 1
}

func main() {
	switch switchexpr2() {
	case case4():
		println("exec case4")
		fallthrough
	case case5():
		println("exec case5")
		fallthrough
	default:
		println("exec default")
	}
}
