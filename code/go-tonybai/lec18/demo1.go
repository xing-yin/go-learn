package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/4
 * @Desc:
 */

func main() {

	if a, c := f(), h(); a > 0 {
		println(a)
	} else if b := f(); b > 0 {
		println(a, b)
	} else {
		println(a, b, c)
	}
}

func h() bool {
	return true
}

func f() int {
	return 2
}
