package main

import (
	"fmt"
	s1 "strings"
)

/**
字符串函数
标准库的 strings 包提供了很多有用的字符串相关的函数。
这里是一些用来让你对这个包有个初步了解的例子。

更多函数参见: https://golang.org/pkg/strings/
*/

// 我们给 fmt.Println 一个短名字的别名，我们随后将会经常用到
var p = fmt.Println

func main() {

	// 这是一些 strings 中的函数例子。注意他们都是包中的函数，不是字符串对象自身的方法，
	// 这意味着我们需要考虑在调用时传递字符作为第一个参数进行传递。
	p("contains: ", s1.Contains("test", "es"))
	p("count: ", s1.Count("tests", "s"))
	p("hasPrefix: ", s1.HasPrefix("test", "st"))
	p("index: ", s1.Index("test", "e"))
	p("join: ", s1.Join([]string{"a", "b", "c"}, "-"))
	p("repeat: ", s1.Repeat("a", 5))
	p("replace: ", s1.Replace("foo", "o", "a", -1))
	p("split: ", s1.Split("a-b-c-d", "-"))
	p("toLower: ", s1.ToLower("TEST"))
	p("toUpper: ", s1.ToUpper("test"))

}
