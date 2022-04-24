package main

import "fmt"

func vals(a int, b int) (int, int) {
	return a, b
}

func main() {

	a, b := vals(1, 2)
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals(3, 4)
	fmt.Println(c)

}
