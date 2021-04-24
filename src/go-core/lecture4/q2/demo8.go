package main

import (
	"flag"
	"fmt"
)

// go 语言的类型推断的好处：程序更加灵活，更易于重构；不会增加维护成本，不会损失运行效率
func main() {
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("hello,%v\n", *name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}

// 将来重构为下面，main不影响
//func getTheFlag() *int {
//	return flag.String("name", 1, "The greeting object.")
//}
