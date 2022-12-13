package main

import "fmt"

func protect(f func()) { //以安全模式来运行所有传入的方法
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f()
}

func main() {
	protect(func() { //模拟函数1
		fmt.Println("This is function 1")
		panic("Serious bug from function 1")
	})
	protect(func() { //模拟函数2
		fmt.Println("This is function 2")
		panic("Serious bug from function 2")
	})
	fmt.Println("valid code")
}
