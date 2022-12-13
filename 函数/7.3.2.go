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
	f1 := addSub
	sum, sub := f1(a, b)
	fmt.Println(a, "+", b, "=", sum)
	fmt.Println(a, "-", b, "=", sub)
}
