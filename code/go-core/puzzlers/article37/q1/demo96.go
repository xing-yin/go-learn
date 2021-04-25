package main

import (
	"errors"
	"fmt"
	"go-learn/code/go-core/puzzlers/article37/common"
	"go-learn/code/go-core/puzzlers/article37/common/op"
	"os"
	"runtime/pprof"
)

var (
	profileName = "cpuprofile.out"
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("CPU profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	if err := startCPUProfile(f); err != nil {
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}
	if err = common.Execute(op.CPUProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	stopCPUProfile()
}

func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}
