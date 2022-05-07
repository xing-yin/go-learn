package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/4
 * @Desc:
 */

func main() {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			firstEven = sl[i]
			break
		}
	}
	println(firstEven)

	var gold = 38
	var sl2 = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outerloop:
	for i := 0; i < len(sl2); i++ {
		for j := 0; j < len(sl2[i]); j++ {
			if sl2[i][j] == gold {
				fmt.Printf("found gold at [%d,%d]\n", i, j)
				break outerloop
			}
		}
	}

}
