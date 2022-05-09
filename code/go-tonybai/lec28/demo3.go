package main

import (
	"fmt"
	"io"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc: 接口类型的方法集合中放入首字母小写的非导出方法也合法
 */

type cancler interface {
	cancle(removeFromParent bool, err error)
	Done() <-chan struct{}
}

// 空接口类型
type EmptyInterface interface {
}

func main() {
	var err error
	var r io.Reader
	fmt.Println(err)
	fmt.Println(r)
}
