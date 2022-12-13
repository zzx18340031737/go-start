package main

import "fmt"

type Book struct {
	title  string
	author string
	num    int
	id     int
}

func main() {
	book1 := &Book{}
	book1.title = "Go语言"
	book1.author = "Tom"
	book1.num = 20
	book1.id = 152368
	fmt.Println("title:", book1.title)
	fmt.Println("author:", book1.author)
	fmt.Println("num:", book1.num)
	fmt.Println("id:", book1.id)
}
