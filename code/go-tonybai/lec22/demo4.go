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

var ErrSentinal = errors.New("the underlying sentinel error")

func main() {
	err1 := fmt.Errorf("wrap sentinel:%w", ErrSentinal)
	err2 := fmt.Errorf("wrap err1:%w", err1)
	println(err2 == ErrSentinal)

	// errors.Is 方法会比较错误链，如果匹配返回true，否则返回false
	if errors.Is(err2, ErrSentinal) {
		println("err2 is ErrSentinal")
		return
	}

	println("err2 is not ErrSentinal")
}
