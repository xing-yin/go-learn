package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Account struct {
	Username string
	Password string
}

func TestReflect(t *testing.T) {
	// reflect.ValueOf() 获取指针对应的反射值
	// reflect.Indirect() 获取指针指向的对象的反射值
	typ := reflect.Indirect(reflect.ValueOf(&Account{})).Type()
	// (reflect.Type).Name() 返回类名(字符串)
	fmt.Println(typ.Name())

	for i := 0; i < typ.NumField(); i++ {
		// (reflect.Type).Field(i) 获取第 i 个成员变量
		field := typ.Field(i)
		fmt.Println(field.Name)
	}
}
