package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 重构：将变量重构为命名常量
const finalWord = "GO!"
const countdownStart = 3

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second}
	Countdown(os.Stdout, sleeper)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}

	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// 记录操作的行为（确保打印顺序按照预期执行）
const sleep = "sleep"
const write = "write"

// CountdownOperationsSpy 同时实现了 io.writer 和 Sleeper，把每一次调用记录到 slice
type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}
