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
	bookBorrow := &BookBorrow{}
	bookNotBorrow := &BookNotBorrow{}
	fmt.Println(bookBorrow)
	fmt.Println(bookNotBorrow)
}
