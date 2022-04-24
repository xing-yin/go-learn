package main

import "fmt"

func fac(num int) int {
	if num == 1 {
		return 1
	}
	return num * fac(num-1)
}

func main() {
	res := fac(7)
	fmt.Println(res)
}
