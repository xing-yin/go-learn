package series

import "fmt"

// init 函数会在 main 被执⾏行行前，所有依赖的 package 的 init ⽅方法都会被执⾏行行
func init() {
	fmt.Println("init1")
}

// 包的每个源⽂文件也可以有多个 init 函数，这点⽐比较特殊
func init() {
	fmt.Println("init2")
}

func square(n int) int {
	return n
}

func Square(n int) int {
	return n * n
}

// 让外面的包可以引用：函数名首字母要大写
func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
