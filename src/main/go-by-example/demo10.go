package main

import "fmt"

/**
关联数组
map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希或者字典 ）
*/
func main() {

	//要创建一个空 map，需要使用内建的 make:make(map[key-type]val-type)
	m := make(map[string]int)

	//使用典型的 make[key] = val 语法来设置键值对
	m["k1"] = 7
	m["k2"] = 13

	// 使用例如 Println 来打印一个 map 将会输出所有的键值对
	fmt.Println("map:", m)

	//使用 name[key] 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//当对一个 map 调用内建的 len 时，返回的是键值对数目
	fmt.Println("len:", len(m))

	//内建的 delete 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	//当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。
	// 这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。
	_, prs := m["k2"]
	//_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//你也可以通过这个语法在同一行申明和初始化一个新的map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}

/**
_, prs := m["k2"]
占位符，意思是那个位置本应赋给某个值，但是咱们不需要这个值。
所以就把该值赋给下划线，意思是丢掉不要。
这样编译器可以更好的优化，任何类型的单个值都可以丢给下划线。
这种情况是占位用的，方法返回两个结果，而你只想要一个结果。
那另一个就用 "_" 占位，而如果用变量的话，不使用，编译器是会报错的。
*/
