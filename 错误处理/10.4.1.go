package main

import "fmt"

func protect() {
	defer func() {
		fmt.Println("func protect exit")
	}()
	panic("Serious bug") //触发宕机
}

func main() { //protec函数退出前执行defer语句
	defer func() {
		fmt.Println("func main exit")
	}()
	protect()
	fmt.Println("Invalid code")
}
