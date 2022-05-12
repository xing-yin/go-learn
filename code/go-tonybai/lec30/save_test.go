package main_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/12
 * @Desc:
 */

func TestSave(t *testing.T) {
	b := make([]byte, 0, 128)
	buf := bytes.NewBuffer(b)
	data := []byte("hello, go")
	err := Save(buf, data)
	if err != nil {
		t.Errorf("want nil,actual %s", err.Error())
	}

	saved := buf.Bytes()
	if !reflect.DeepEqual(saved, data) {
		t.Errorf("want %s,actual %s", string(data), string(saved))
	}
}

func Save(w io.Writer, data []byte) error {
	w.Write(data)
	return nil
}
