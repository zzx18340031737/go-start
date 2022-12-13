package main

import "fmt"

func main() {
	var student = []string{"Tom", "Ben", "Peter", "Danny"}
	for k, v := range student {
		fmt.Println("切片下标：", k, ",对应元素", v)
	}
}
