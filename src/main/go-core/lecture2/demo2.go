package main

import (
	// 需在此处添加代码。[1]
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 需在此处添加代码。[3]
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
