package main

// 需求:自定义命令源码文件的参数使用说明

import (
	"flag"
	"fmt"
	"os"
)

var name2 string

func init() {
	flag.StringVar(&name2, "name2", "everyone", "The greeting object")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "question")
	}
	flag.PrintDefaults()
	// 真正解析命令参数，并把它们的值赋给相应的变量
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name2)
}

// 运行说明：go run demo2.go -name3="yx"
// go run demo2.go --help 查看该命令源码文件的参数说明
