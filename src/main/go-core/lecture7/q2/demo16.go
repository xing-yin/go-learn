package main

import "fmt"

func main() {
	// 示例1
	s1 := make([]int, 0)
	fmt.Printf("the capacity of s1:%d\n", cap(s1))
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
		fmt.Printf("the element(%d): len(%d): cap:(%d)\n", i, len(s1), cap(s1))
	}
	fmt.Println()

	// 示例2
	s2 := make([]int, 1024)
	fmt.Printf("the capacity of s2:%d\n", cap(s2))
	s2e1 := append(s2, make([]int, 200)...)
	fmt.Printf("s2e1: len: %d, cap: %d\n", len(s2e1), cap(s2e1))
	s2e2 := append(s2, make([]int, 400)...)
	fmt.Printf("s2e2: len: %d, cap: %d\n", len(s2e2), cap(s2e2))
	s2e3 := append(s2, make([]int, 600)...)
	fmt.Printf("s2e3: len: %d, cap: %d\n", len(s2e3), cap(s2e3))
	fmt.Println()

	// 示例3
	s3 := make([]int, 10)
	fmt.Printf("the capacity of s3:%d\n", cap(s3))
	s3a := append(s3, make([]int, 11)...)
	fmt.Printf("s3a: len: %d, cap: %d\n", len(s3a), cap(s3a))
	s3b := append(s3a, make([]int, 23)...)
	fmt.Printf("s3b: len: %d, cap: %d\n", len(s3b), cap(s3b))
	s3c := append(s3b, make([]int, 45)...)
	fmt.Printf("s3c: len: %d, cap: %d\n", len(s3c), cap(s3c))

}
