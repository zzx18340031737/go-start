package main

import "fmt"

type Book struct {
	title  string
	author string
	num    int
	id     int
}

type BookBorrow struct {
	Book
	borrowTime string
}

type BookNotBorrow struct {
	Book
	readTime string
}

func main() {
	bookBorrow := &BookBorrow{
		Book: Book{
			"Go语言",
			"Tom",
			20,
			152368,
		},
		borrowTime: "30",
	}
	fmt.Println(bookBorrow)
	bookNotBorrow := &BookNotBorrow{}
	bookNotBorrow.title = "Python语言"
	bookNotBorrow.author = "Peter"
	bookNotBorrow.num = 10
	bookNotBorrow.id = 152369
	bookNotBorrow.readTime = "50"
	fmt.Println(bookNotBorrow)
}
