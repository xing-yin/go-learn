package main

import (
	"flag"
	"go-learn/code/go-core/lecture3/q2/lib"
	"go-learn/code/go-core/lecture3/q4/internal"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	lib.Hello(name)
	// 是 internal 的父包及其子包，有访问权限
	internal.Hello(os.Stdout, name)
}
