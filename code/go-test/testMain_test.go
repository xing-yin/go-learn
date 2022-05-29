package test

import (
	"fmt"
	"testing"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/20
 * @Desc:
 */

func TestMain(m *testing.M) {
	fmt.Println("do some setup")
	m.Run()
	fmt.Println("do some cleanup")
}
