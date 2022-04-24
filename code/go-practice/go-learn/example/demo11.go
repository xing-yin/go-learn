package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num > 0 {
			fmt.Println("index is:", i)
		}
	}

	kvs := map[string]int{"k1": 1, "k2": 2}
	for key, value := range kvs {
		fmt.Printf("% s -> %d\n", key, value)
	}

	for i, c := range "gogo" {
		fmt.Println(i, c)
	}
}
