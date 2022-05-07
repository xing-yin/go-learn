package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

func case1() int {
	println("case1 expr")
	return 1
}

func case2_1() int {
	println("case2_1 expr")
	return 2
}

func case2_2() int {
	println("case2_2 expr")
	return 3
}

func case3() int {
	println("case3 expr")
	return 4
}

func switchexpr() int {
	println("switch expr")
	return 2
}

func main() {
	switch switchexpr() {
	default:
		println("exec default")
	case case1():
		println("exec case1")
	case case2_1(), case2_2():
		println("exec case2")
	case case3():
		println("exec case3")
	}
}
