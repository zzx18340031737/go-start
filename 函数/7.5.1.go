package main

import "fmt"

func main() {
	func(data string) {
		fmt.Println("Hello " + data)
	}("world!")
}
