package main

import "flag"

/**
把命令源码文件中的代码拆分到其他库源码文件
*/

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello(name)
}
