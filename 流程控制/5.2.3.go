package main

import "fmt"

func main() {
	i := 1
OuterLook:
	for {
		for {
			if i > 5 {
				break OuterLook //跳出outerLook标签对应的循环
			}
			fmt.Println(i)
			i++
		}
	}
}
