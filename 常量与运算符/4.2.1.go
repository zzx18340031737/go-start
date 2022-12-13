package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("a/b=", a/b)
	fmt.Println("a%b=", a%b)
	a++
	fmt.Println("a++后的值为：", a)
	b--
	fmt.Println("b--后的值为：", b)

}
