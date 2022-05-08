package main

import "time"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */
func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

func main() {
	teardown := setup("demo")
	defer teardown()
	println("do some business stuff")

	time.AfterFunc(time.Second*2, func() {
		println("timer fired")
	})
}
