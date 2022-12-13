package main

import "fmt"

func main() {
	panic("Serious bug")
	fmt.Println("Invalid code") //程序退出 ，无法执行该代码
}
