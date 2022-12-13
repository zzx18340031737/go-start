package main

import "fmt"

type Book struct {
	title  string
	anthor string
	num    int
	id     int
}

func main() {
	book1 := &Book{}
	fmt.Println(book1)
}
