package main

import "fmt"

func main() {
	a := 101
	if a > 100 {
		fmt.Println(a, " > 100")
	} else if a == 100 {
		fmt.Println(a, " = 100")
	} else {
		fmt.Println(a, " < 100")
	}
}
