package main

import "fmt"

func main() {
	switch 1 + 1 {
	case 1:
		fmt.Println("1+1=1")
	case 2:
		fmt.Println("1+1=2")
	case 3:
		fmt.Println("1+1=3")
	default:
		fmt.Println("1+1不等于1或2或3")
	}
}
