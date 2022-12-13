package main

import "fmt"

func main() {
	var student = [...]string{"Tom", "Ben", "Peter"}
	for k, v := range student {
		fmt.Println("数组下标:", k, ",对应元素", v)
	}
}
