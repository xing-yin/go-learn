package main

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func TestName(t *testing.T) {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
