package main

import "fmt"

type BookBorrow struct {
	Book struct { //内嵌匿名结构体
		title  string
		author string
		num    int
		id     int
	}
	borrowTime string
}

func main() {
	bookBorrow := &BookBorrow{
		Book: struct { //声明类型
			title  string
			author string
			num    int
			id     int
		}{
			"Go语言",
			"Tom",
			20,
			152368,
		},
		borrowTime: "30",
	}
	fmt.Println(bookBorrow)
}
