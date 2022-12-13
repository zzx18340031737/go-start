package main

import "fmt"

func addall(slice ...int) {
	sum := 0
	for _, value := range slice {
		sum = sum + value
	}
	fmt.Println(sum)
}

func add(num ...int) {
	addall(num...)
}

func main() {
	add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
