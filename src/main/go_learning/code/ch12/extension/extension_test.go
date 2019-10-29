package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 使用复合的方式代替继承
type Dog struct {
	// 使用指针不会调用"子类"的方法("Wang!")，需要Dog重新实现SpeakTo()
	//p *Pet
	// go 中的匿名类型嵌⼊
	Pet
}

// 子类并不是真正继承了⽗类的方法
func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	// go 中不⽀支持子类替换
	//var dog Pet := new(Dog)
	dog := new(Dog)
	// 强制转换也不支持
	//var p = (* Pet)(dog)
	// 子类并不是真正继承了⽗类的方法，不会打印 "wang!"
	dog.SpeakTo("Chao")
}
