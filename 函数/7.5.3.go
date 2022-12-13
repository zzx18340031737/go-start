package main

import "fmt"

func main() {
	num := 1
	fmt.Printf("%p\n", &num)
	func() {
		num++
		fmt.Println(num)
		fmt.Printf("%p\n", &num)
	}()
	func() {
		num++
		fmt.Println(num)
		fmt.Printf("%p\n", &num)
	}()
}
