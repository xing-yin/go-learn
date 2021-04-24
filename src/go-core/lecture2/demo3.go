package main

import (
	"flag"
	"fmt"
	"os"
)

/**
自定义命令源码文件的参数使用说明
*/
var name1 string

func init() {
	flag.StringVar(&name1, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name1)
}
