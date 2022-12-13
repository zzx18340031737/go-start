package main

import "fmt"

func main() {
	i := 1
	for {
		for {
			if i > 5 {
				fmt.Println("跳出内层for循环")
				break
			}
			fmt.Println(i)
			i++
		}
		fmt.Println("跳出外层for循环")
		break
	}
}
