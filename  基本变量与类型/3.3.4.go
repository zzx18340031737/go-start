package main

import "fmt"

func main() {
	var p *int
	fmt.Println("指针变量p指向的地址为:", p)
	fmt.Println("指针变量p所指向的内容为:", *p)
}
