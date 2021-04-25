package main

import (
	"flag"
	"fmt"
)

func main() {
	// 声明变量方式1
	//var name string
	//flag.StringVar(&name,"name","everyone","The greeting object.")

	// 方式2
	//var name = flag.String("name", "everyone", "The greeting object.")

	// 方式3
	name := flag.String("name", "everyone", "The greeting object.")

	flag.Parse()
	fmt.Printf("hello,%v\n", *name)

	// 适用于方式1
	//flag.Parse()
	//fmt.Printf("hello,%v\n",name)
}
