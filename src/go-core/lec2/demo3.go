package main

// flag.CommandLine 示例演示

import (
	"flag"
	"fmt"
	"os"
)

var name3 string

func init() {
	flag.StringVar(&name3, "name3", "everyone", "The greeting object")
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "question")
		flag.PrintDefaults()
	}
}

func main() {
	// 真正解析命令参数，并把它们的值赋给相应的变量
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name3)
}

// 运行说明：go run demo.go -name3="yx"
// go run demo3.go --help 查看该命令源码文件的参数说明
