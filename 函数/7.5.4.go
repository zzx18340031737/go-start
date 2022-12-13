package main

import "fmt"

func addOne(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	a1 := addOne(0)
	fmt.Println(a1()) //0+1=1
	fmt.Println(a1()) //1+1=2
	a2 := addOne(10)
	fmt.Println(a2())
	fmt.Println("a1闭包的地址为：")
	fmt.Printf("%p\n", &a1)
	fmt.Println("a2闭包的地址为：")
	fmt.Printf("%p\n", &a2)
}
