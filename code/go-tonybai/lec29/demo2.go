package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/11
 * @Desc:
 */

type QuackableAnimal interface {
	Quack()
}

type Duck struct {
}

func (d Duck) Quack() {
	println("duck quack")
}

type Dog struct {
}

func (d Dog) Quack() {
	println("dog quack")
}

type Bird struct {
}

func (b Bird) Quack() {
	println("bird quack")
}

func AnimalQuackInForest(a QuackableAnimal) {
	a.Quack()
}

func main() {
	animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
	for _, animal := range animals {
		AnimalQuackInForest(animal)
	}
}
