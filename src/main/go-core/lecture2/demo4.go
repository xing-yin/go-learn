package main

import (
	"flag"
	"fmt"
	"os"
)

/**
自定义命令源码文件的参数使用说明
*/
var name2 string

func init() {
	flag.StringVar(&name2, "name2", "everyone", "The greeting object.")
}

func main() {
	//flag.CommandLine = flag.NewFlagSet("",flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name2)
}
