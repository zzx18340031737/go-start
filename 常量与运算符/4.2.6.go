package main

import "fmt"

func main() {
	a := 1
	var p *int //声明指针p，指针类型为int类型
	p = &a     //指针p取变量a的地址
	fmt.Println("^a = ", ^a)
	fmt.Println("变量a的地址为：", p)

}
