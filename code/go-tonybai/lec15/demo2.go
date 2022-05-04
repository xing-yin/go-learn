package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/1
 * @Desc:
 */

func main() {

	//var nums = []int {1,2,3,4,5,6}
	//fmt.Println(len(nums))
	//nums = append(nums, 7)
	//fmt.Println(len(nums))

	//sl1 := make([]byte, 6, 10)
	//sl2 := make([]byte, 6)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl3 := arr[3:7:9]
	fmt.Println(sl3)

	sl4 := arr[3:7]
	fmt.Println(sl4)

	sl3[0] += 10
	fmt.Println(arr[3])

}
