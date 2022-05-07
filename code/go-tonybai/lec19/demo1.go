package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/4
 * @Desc:
 */

func main() {

	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for i, j, k := 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k = i+1, j+1, k+5 {
		sum += (i + j + k)
		println(sum)
	}

	for i := 0; i < 10; {
		i++
	}

	i1 := 0
	for ; i1 < 5; i1++ {
		println(i1)
	}

	i2 := 0
	for i2 < 5 {
		i2++
	}

	for {

	}
}
