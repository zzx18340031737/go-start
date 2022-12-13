package main

import "fmt"

func main() {
	fmt.Println("Hello")
	goto sign
	fmt.Println("无效代码")
sign:
	fmt.Println("world")
}
