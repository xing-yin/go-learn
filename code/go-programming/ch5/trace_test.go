package ch5

import (
	"log"
	"testing"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")

	// lots of work
	time.Sleep(10 * time.Second) // simulate slow
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func TestTrace(t *testing.T) {
	bigSlowOperation()
}
