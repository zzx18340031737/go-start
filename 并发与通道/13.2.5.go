package main

import "fmt"

func main() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("main")
	}
}
