package basic

import (
	"fmt"
	"testing"
)

//func main() {
//
//	/// 错误类型
//	err := errors.New("error msg")
//	if err != nil {
//		fmt.Print(err)
//	}
//
//}

/// ==========================一些技巧
// 1.分组声明
// 优化前
const i01 = 100
const pi01 = 3.14
const prefix01 = "Go_"

var i11 int
var pi11 float32
var prefix11 string

// 优化后
const (
	i02      = 100
	pi02     = 3.14
	prefix02 = "Go_"
)

var (
	i12      int
	pi12     float32
	prefix13 string
)

// 2.iota枚举: 这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1：
const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota // a=0
	b       = "B"
	c       = iota             // c=2
	d, e, f = iota, iota, iota // d,e,f 均为 3
	g       = iota             //  g=4
)

// Go程序设计的一些规则
//Go之所以会那么简洁，是因为它有一些默认的行为：
// - 大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
// - 大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

func Test1(t *testing.T) {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}

/// todo ========================== array、slice、map
// 1.array

func TestArray(t *testing.T) {
	var arr [10]int // 声明了一个int类型的数组
	arr[0] = 1      // 数组下标是从0开始的
	arr[1] = 2
	fmt.Printf("the first element is %d\n", arr[0])
	fmt.Printf("the last element is %d\n", arr[9])

	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 数组里面的值还是数组，能实现吗？当然咯，Go支持嵌套数组，即多维数组。比如下面的代码就声明了一个二维数组：
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray)
	fmt.Println(easyArray)
}

// 2.slice
// 在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice。
// slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

// 和声明array一样，只是少了长度
var fslice []int

func TestSlice(t *testing.T) {

	// 01 声明一个slice，并初始化数据
	slice := []byte{'a', 'b', 'c'}
	fmt.Println(slice)

	// 02 slice可以从一个数组或一个已经存在的slice中再次声明。
	// slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j](左闭右开)，它的长度是j-i。
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// 声明 2 个含有byte的slice
	var a, b, c []byte

	// a 指向数组的第3个元素开始，到第5个元素结束
	a = ar[2:5] //现在a含有的元素: ar[2]、ar[3]和ar[4]

	// b 是数组的另一个 slice
	b = ar[3:5] //现在b含有的元素: ar[3]、ar[4]

	// c 是从 slice 中取 slice
	c = a[0:2]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 03 slice有一些简便的操作
	//- slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	//- slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	//- 如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]

	// 04 slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值

	// 05从Go1.2开始slice支持了三个参数的slice，之前我们一直采用这种方式在slice或者array基础上来获取一个slice
	var array [10]int
	slice2 := array[2:4]
	fmt.Println("05====")
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// 可以指定这个容量（7-2=5）
	slice3 := array[2:4:7]
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}

// ⚠️注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符

/// todo ========================== map
// map 就是Python中字典的概念，它的格式为map[keyType]valueType
// map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int｀类型，
// 而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。

// 使用map过程中需要注意的几点：
//- map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
//- map的长度是不固定的，也就是和slice一样，也是一种引用类型
//- 内置的len函数同样适用于map，返回map拥有的key的数量
//- map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
//- map和其他基本型别不同，它不是 thread-safe ，在多个go-routine存取时，必须使用mutex lock机制

func TestMap(t *testing.T) {
	// 01 声明
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string]int
	//// 另一种map的声明方式
	//numbers2 := make(map[string]int)
	//
	//numbers["one"] = 1
	//numbers["two"] = 2
	//numbers2["one"] = 1
	//fmt.Println("the two element is:", numbers["two"])
	//fmt.Println("the one element is:", numbers2["one"])

	// 02 初始化一个字典
	rating := map[string]float32{"c": 5, "go": 4.5, "python": 4.5, "java": 4}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["c#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	// 删除 key 为 c 的元素
	delete(rating, "c")
	fmt.Println(rating)

	// 03 map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变：
	m := make(map[string]string)
	m["hello"] = "go"
	fmt.Println("before m:", m)
	m1 := m
	m1["hello"] = "world" // 现在m["hello"]的值已经是world了
	fmt.Println("after m:", m)
}

/// 零值

//“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”

//int     0
//int8    0
//int32   0
//int64   0
//uint    0x0
//rune    0 //rune的实际类型是 int32
//byte    0x0 // byte的实际类型是 uint8
//float32 0 //长度为 4 byte
//float64 0 //长度为 8 byte
//bool    false
//string  ""
