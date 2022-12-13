package main

import "fmt"

func protect() {
	defer func() {
		if err := recover(); err != nil { //recover()获取panic（）传入的参数
			fmt.Println(err)
		}
	}()
	panic("Serious bug")
}

func main() {
	protect()
	fmt.Println("valid code")
}
