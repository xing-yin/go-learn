package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test705(t *testing.T) {
	var w io.Writer
	w = nil
	fmt.Printf("%T\n", w)
	// w.Write([]byte("hello")) // panic: runtime error: invalid memory address or nil pointer dereference

	w = os.Stdout
	fmt.Printf("%T\n", w)
	fmt.Println(w.Write([]byte("hello")))

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	fmt.Println(w.Write([]byte("hello")))
}
