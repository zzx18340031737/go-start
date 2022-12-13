package main

import "fmt"

type Book struct {
	title  string
	author string
	num    int
	id     int
}

func main() {
	book1 := &Book{
		title:  "Go语言",
		author: "Tom",
		num:    20,
		id:     152368,
	}
	fmt.Println("title:", book1.title)
	fmt.Println("author:", book1.author)
	fmt.Println("num:", book1.num)
	fmt.Println("id:", book1.id)
}
