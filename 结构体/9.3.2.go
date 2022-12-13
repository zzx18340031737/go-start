package main

import "fmt"

type Book struct {
	title  string
	author string
	num    int
	id     int
}

func main() {
	book1 := new(Book)
	fmt.Println(book1)
}
