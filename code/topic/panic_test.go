package main

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	fmt.Println("before panic")
	panic("crash")
	fmt.Println("after panic")
}

func TestPanic2(t *testing.T) {
	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
}

func TestRecover(t *testing.T) {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}

func TestInvokeRecover(t *testing.T) {
	TestRecover(t)
	fmt.Println("after recover")
}
