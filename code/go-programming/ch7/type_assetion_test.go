package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAssertion(t *testing.T) {
	//func1()

	//func2()

	func3()
}

func func3() {
	var w io.Writer
	w = os.Stdout
	f, ok := w.(*os.File)
	if !ok {
		fmt.Errorf("failed to os.File")
	}
	b, ok := w.(*bytes.Buffer)
	if !ok {
		fmt.Errorf("failed to bytes.Buffer")
	}
	fmt.Println(f, b)
}

func func2() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	w = new(ByteCounter)
	rw = w.(io.ReadWriter) // panic: interface conversion: *main.ByteCounter is not io.ReadWriter: missing method Read
	fmt.Println(rw)
}

func func1() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)      // success: f == os.Stdout
	c := w.(*bytes.Buffer) // panic: io.Writer is *os.File, not *bytes.Buffer
	fmt.Println(f, c)
}
