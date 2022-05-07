package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/4
 * @Desc:
 */

func main() {
	//var m = []int{1, 2, 3, 4, 5}

	// 错误写法
	//for i, v := range m {
	//	go func() {
	//		time.Sleep(time.Second * 3)
	//		fmt.Println(i, v)
	//	}()
	//}

	// 正确写法
	//for i, v := range m {
	//	go func(i, v int) {
	//		time.Sleep(time.Second * 3)
	//		fmt.Println(i, v)
	//	}(i, v)
	//}
	//
	//time.Sleep(time.Second * 10)
}
