package main

import (
	"bytes"
	"fmt"
	"testing"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(c)
}

func TestEmptyInterface(t *testing.T) {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)
}
