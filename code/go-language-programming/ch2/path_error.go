package ch2

import (
	"os"
	"syscall"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e PathError) Error() string {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

func Stat(name string) (fi os.FileInfo, err error) {
	var stat syscall.Stat_t

	err = syscall.Stat(name, &stat)
	if err != nil {
		return nil, &PathError{"stat", name, err}
	}
	return nil, nil
}
