package _interface2

import (
	"fmt"
)

/// todo=================================什么是interface
// interface是一组method签名的组合，我们通过interface来定义对象的一组行为。

/// todo=================================interface类型
// interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human对象实现SayHi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi,name:%s,phone:%s", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("la la la ...", lyrics)
}

// Human对象实现Guzzle(狂欢)方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的SayHi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

// Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

// todo 定义 interface
// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这3个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

// 上面的Men interface 被Human、Student和Employee实现。
// 同理，一个对象可以实现任意多个interface，例如上面的Student实现了Men和YoungChap两个interface。

// 任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。
