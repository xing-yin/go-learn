package interface_value

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

/// todo=================================interface值
// 那么interface里面到底能存什么值呢？
// 如果我们定义了一个 interface 的变量，那么这个变量里面可以存实现这个interface的任意类型的对象。
// 例如上面例子中，我们定义了一个Men interface类型的变量m，那么m里面可以存Human、Student或者Employee值。

//因为m能够持有这三种类型的对象，所以我们可以定义一个包含Men类型元素的slice，
//这个slice可以被赋予实现了Men接口的任意结构的对象，这个和我们传统意义上面的slice有所不同。

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func TestInterfaceValue(t *testing.T) {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}

// 通过上面的代码，你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
// Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。

/// todo=================================空interface
// 空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。
// 空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，
// 因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。

// 定义 a 为空接口
var a interface{}
var i int = 5
var s string = "Hello world"

// a可以存储任意类型的数值
func TestEmptyInterface(t *testing.T) {
	a = i
	a = s
}

// 一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。

/// todo=================================interface函数参数
// interface的变量可以持有任意实现该interface类型的对象，这给我们编写函数(包括method)提供了一些额外的思考，我们是不是可以通过定义interface参数，让函数接受各种类型的参数。

// 01
// 举个例子：fmt.Println是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，你会看到这样一个定义:
//type Stringer interface {
//     String() string
//}
// 任何实现了String方法的类型都能作为参数被fmt.Println调用
type Human1 struct {
	name  string
	age   int
	phone string
}

// 通过这个方法 human 实现了 fmt.Stringer
func (h Human1) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func TestParam(t *testing.T) {
	Bob := Human1{"Bob", 39, "432423"}
	fmt.Println("this human is:", Bob)
}

/// todo=================================interface变量存储的类型
// 我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法

// 01Comma-ok断言
// Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，
// 这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型
// 如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false

type Element interface{}

type List []Element

type Person struct {
	name string
	age  int
}

// 定义了 String 方法，实现了 fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " -age: " + strconv.Itoa(p.age) + " years"
}

func Test1(t *testing.T) {
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "hello" // a string
	list[2] = Person{"Denny", 20}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is an int and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is an int and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type", index)
		}
	}
}

// 02switch测试
// 用switch 实现01代码
func Test2(t *testing.T) {
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "hello" // a string
	list[2] = Person{"Denny", 20}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is an int and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is an int and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type", index)
		}
	}
}

// ⚠️有一点需要强调的是：element.(type)语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用comma-ok

/// todo=================================嵌入interface
// Go里面真正吸引人的是它内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，
// 那么相同的逻辑引入到interface里面，那不是更加完美了。
// 如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。

// 可以看到源码包container/heap里面有这样的一个定义
type Interface interface {
	sort.Interface      //嵌入字段sort.Interface
	Push(x interface{}) //a Push method to push elements into the heap
	Pop() interface{}   //a Pop elements that pops elements from the heap
}

// 另一个例子就是io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个 interface：

/// todo=================================反射
// Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。
// 我们一般用到的包是reflect包。如何运用reflect包，官方的这篇文章详细的讲解了reflect包的实现原理，laws of reflection:https://blog.golang.org/laws-of-reflection

//使用reflect一般分成三步，下面简要的讲解一下：要去反射是一个类型的值(这些值都实现了空interface)，
// 首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。
// 这两种获取方式如下：

func TestReflect(t *testing.T) {
	//t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	//v2 := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	//
	//// 转化为reflect对象之后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值，例如
	//tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
	//name := v2.Elem().Field(0).String()  //获取存储在第一个字段里面的值
	//
	//// 获取反射值能返回相应的类型和数值
	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//fmt.Println("value:", v.Float())

	// 最后，反射的话，那么反射的字段必须是可修改的，我们前面学习过传值和传引用，
	// 这个里面也是一样的道理。反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误
	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//v.SetFloat(7.1)

	// 如果要修改相应的值，必须这样写
	//var x float64 = 3.4
	//p := reflect.ValueOf(&x)
	//v := p.Elem()
	//v.SetFloat(7.1)
}
