package main

import "errors"

// ErrNotFound 重构：摆脱魔法值
var ErrNotFound = errors.New("could not found the world you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}
