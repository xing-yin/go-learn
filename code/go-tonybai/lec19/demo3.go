package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/4
 * @Desc:
 */

func main() {
	var sum int
	var sl = []int{1, 2, 3, 4, 5, 6}

loop:
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			continue loop
		}
		sum += sl[i]
	}
	println(sum)

	var sl2 = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outerloop:
	for i := 0; i < len(sl2); i++ {
		for j := 0; j < len(sl2[i]); j++ {
			if sl2[i][j] == 13 {
				fmt.Printf("found 13 at [%d,%d]\n", i, j)
				continue outerloop
			}
		}
	}
}
