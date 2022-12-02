package main

import "fmt"

func main() {
	var c int
	a := 1
	b := 2
	c = a
	a = b
	b = c
	fmt.Println(a)
	fmt.Println(b)
}
