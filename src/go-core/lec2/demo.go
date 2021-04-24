package main

// 需求:在注释处添加相应的代码，并让程序实现”根据运行程序时给定的参数问候某人”的功能

import (
	"flag"
	"fmt" // 需在此处添加代码。[1]
)

var name string

func init() {
	// 需在此处添加代码。[2]
	flag.StringVar(&name, "name", "everyone", "The greeting object")
	// 相同功能的方法（直接返回一个已经分配好的用于存储命令参数值的地址）
	//flag.String("name4","everyone","The greeting object")
}

func main() {
	// 需在此处添加代码。[3]
	// 真正解析命令参数，并把它们的值赋给相应的变量
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

// 运行说明：go run demo.go -name="yx"
// go run demo.go --help 查看该命令源码文件的参数说明
