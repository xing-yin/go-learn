package main

import (
	"errors"
	"fmt"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

type MyError struct {
	s string
}

func (e *MyError) Error() string {
	return e.s
}

func main() {
	var err = &MyError{"MyError error demo"}
	err1 := fmt.Errorf("wrap err:%w", err)
	err2 := fmt.Errorf("wrap err1:%w", err1)
	var e *MyError
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		println(e == err)
		return
	}

	println("MyError is not on the chain of err2")
}
