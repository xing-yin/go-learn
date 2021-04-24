package main

// flag.CommandLine 示例演示

import (
	"flag"
	"os"
)

var name5 string

func init() {
	/// todo 区别于 demo4
	var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)
	cmdLine.StringVar(&name5, "name5", "everyone", "The greeting object")
	cmdLine.Parse(os.Args[1:])
}

func main() {
	// 真正解析命令参数，并把它们的值赋给相应的变量
}

// 运行说明：go run demo5.go -name5="yx"
// go run demo5.go --help 查看该命令源码文件的参数说明
