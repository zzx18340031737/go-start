package main

import "fmt"

type Book struct {
	title  string
	author string
	num    int
	id     int
}

func main() {
	var book1 Book
	fmt.Println(book1)
}
