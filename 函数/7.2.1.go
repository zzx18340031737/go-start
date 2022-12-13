package main

import "fmt"

func addSub(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func main() {
	a := 1
	b := 2
	sum, sub := addSub(a, b) //调用addSub函数
	fmt.Println(a, "+", b, "=", sum)
	fmt.Println(a, "-", b, "=", sub)
}
