package lib

import "fmt"

func Hello(name string) {
	// 是 internal 的父包及其子包，有访问权限
	fmt.Printf("hello,%s\n", name)
}
