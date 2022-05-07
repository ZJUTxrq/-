package main

import "fmt"

type Book struct {
	title string
	auth  string
}

//改变不了lisi
func bookUpdate(book Book) {
	book.auth = "lisi"
}

//改变lisi，必须以指针方式
func bookUp(book *Book) {
	book.auth = "lisi"
}

func main() {
	fmt.Println("hello world")

	var book Book
	book.title = "Goland"
	book.auth = "zhangsan"

	fmt.Println(book)

	bookUpdate(book)
	fmt.Println(book)

	bookUp(&book)
	fmt.Println(book)

	branch()
}

func branch() {
	fmt.Println("分支Test")
}
