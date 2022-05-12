package main

import (
	"errors"
	"fmt"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/11
 * @Desc:
 */

func main() {
	//var err error = 1

	var err error
	err = errors.New("error1")
	fmt.Printf("%T\n", err)
}
