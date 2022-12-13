package main

import "fmt"

func main() {
	var num = [...]int{1, 2, 3, 4}
	for k, v := range num {
		fmt.Println("变量k:", k, " ", "变量v:", v)
	}
}
