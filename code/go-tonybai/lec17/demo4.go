package lec17

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/2
 * @Desc:
 */

type Person struct {
	Name  string
	Phone int
	Add   string
}

type Book2 struct {
	Title  string
	Author Person
}

type Book3 struct {
	Title string
	Person
}

func main() {
	var book Book2
	fmt.Println(book.Author.Phone)

	var book3 Book3
	fmt.Println(book3.Phone)
}
