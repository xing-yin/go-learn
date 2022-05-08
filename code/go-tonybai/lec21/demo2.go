package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

var (
	myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, a...)
	}
)

func main() {
	fmt.Printf("%T\n", myFprintf)            // func(io.Writer, string, ...interface {}) (int, error)
	myFprintf(os.Stdout, "%s\n", "hello,Go") // hello,Go
}
