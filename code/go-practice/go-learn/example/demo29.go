package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/4/25
 * @Desc:
 */

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message receive")
	}

	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
