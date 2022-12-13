package main

import "fmt"

func main() {
	f1 := func(data string) {
		fmt.Println("Hello " + data)
	}
	f1("world!")
}
