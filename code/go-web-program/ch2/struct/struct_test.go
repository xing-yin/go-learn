package _struct

import (
	"fmt"
	"testing"
)

/// todo=================================struct
// Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。

// 01 声明一个 struct 类型
// 例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。如下代码所示:
type person struct {
	name string
	age  int
}

// 02 如何使用struct呢？
var P person // P现在就是person类型的变量了

func TestStruct1(t *testing.T) {
	P.name = "alan"
	P.age = 18
	fmt.Println(P.name)

	// 除了上面这种P的声明使用之外，还有另外几种声明使用方式：
	// 1.按照顺序提供初始化值
	P1 := person{"alan", 18}
	// 2.通过field:value的方式初始化，这样可以任意顺序
	P2 := person{name: "alan", age: 18}
	// 3.当然也可以通过new函数分配一个指针，此处P的类型为 *person
	P3 := new(person)
	fmt.Println(P1, P2, P3)

}

// 03 下面我们看一个完整的使用struct的例子
// 比较 2个人的年龄，返回年龄大的那个人并返回年龄差
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func testStruct2() {
	var tom person
	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	alan := person{name: "alan", age: 18}

	// 按照struct定义顺序初始化值
	bob := person{"bob", 21}

	ta_Older, ta_diff := Older(tom, alan)
	tb_Older, tb_diff := Older(tom, bob)
	ab_Olde, ab_diff := Older(alan, bob)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, ta_Older.name, ta_diff)
	fmt.Println(tb_Older, tb_diff)
	fmt.Println(ab_Olde, ab_diff)
}

/// todo=================================struct的匿名字段
// 我们上面介绍了如何定义一个struct，定义的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

// 01
// 当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
type Human struct {
	name   string
	age    int
	weight int
}

// Student组合了Human struct和string基本类型
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func Test(t *testing.T) {
	// 初始化一个学生
	mark := Student{Human{"mark", 25, 130}, "cs"}

	// 访问相应的字段
	fmt.Println("name is:", mark.name)
	fmt.Println("age is:", mark.age)
	fmt.Println("weight is:", mark.weight)
	fmt.Println("speciality is:", mark.speciality)

	// 修改对应的字段
	mark.speciality = "AI"
	fmt.Println("mark 's speciality is change")
	fmt.Println("speciality is:", mark.speciality)

	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}

// 02 我们看到Student访问属性age和name的时候，就像访问自己所有用的字段一样，对，匿名字段就是这样，能够实现字段的继承。
// 还有比这个更酷的呢，那就是student还能访问Human这个字段作为字段名。
func Test2(t *testing.T) {
	mark := Student{Human{"mark", 25, 130}, "cs"}

	mark.Human = Human{"mark", 20, 100}
	mark.Human.age -= 1
}

// 03 通过匿名访问和修改字段相当的有用，但是不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段的。请看下面的例子

type Skills []string

type Student2 struct {
	Human      // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func Test3(t *testing.T) {
	jane := Student2{Human: Human{"Jane", 18, 100}, speciality: "history"}
	// 访问相应的字段
	fmt.Println("name is:", jane.name)
	fmt.Println("age is:", jane.age)
	fmt.Println("weight is:", jane.weight)
	fmt.Println("speciality is:", jane.speciality)

	// 我们来修改他的skill技能字段
	jane.Skills = []string{"run"}
	fmt.Println("Skills is:", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("now Skills is:", jane.Skills)

	// 修改匿名内置类型的字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

// 04 这里有一个问题：如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？
// Go里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段。
// 这就允许我们去重载通过匿名字段继承的一些字段，当然如果我们想访问重载后对应匿名类型里面的字段，可以通过匿名字段名来访问。

type Human3 struct {
	name  string
	age   int
	phone string
}

type Student3 struct {
	Human3
	phone      string
	speciality string
}

func Test4(t *testing.T) {
	bob := Student3{Human3{"Bob", 23, "6666666"}, "777777777", "math"}
	fmt.Println("bob 's phone is :", bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("bob 's Human phone is :", bob.Human3.phone)
}
