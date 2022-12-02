package main

import "fmt"

func main() {
	num := 1
	var p *int
	p = &num

	fmt.Println("指针变量p所指向的内容：",*p)
	*p = 10
	fmt.Println("指针变量p所指向的内容：",*p)
}
