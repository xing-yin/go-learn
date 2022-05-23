package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/19
 * @Desc:
 */

// Bird 基类
type Bird struct {
	Type string
}

// 鸟的种类
func (bird Bird) Class() string {
	return bird.Type
}

// -------------------------------------------------------

// Birds 定义一个鸟类接口
type Birds interface {
	Name() string
	Class() string
}

// -------------------------------------------------------

// 具体实现"类"1：金丝雀
type Canary struct {
	Bird // 匿名字段
	name string
}

func (c *Canary) Name() string {
	return c.name
}

// 具体实现"类"2：乌鸦
type Crow struct {
	Bird
	name string
}

func (c Crow) Name() string {
	return c.name
}

func NewCanary(name string) *Canary {
	return &Canary{
		Bird: Bird{
			Type: "Canary",
		},
		name: name,
	}
}

func NewCrow(name string) *Crow {
	return &Crow{
		Bird: Bird{
			Type: "Crow",
		},
		name: name,
	}
}

func BirdInfo(birds Birds) {
	fmt.Printf("I'm %s, I belong to %s bird class!\n", birds.Name(), birds.Class())
}

func main() {
	canary := NewCanary("Canary1")
	crow := NewCrow("Crow1")
	BirdInfo(canary)
	BirdInfo(crow)
}
