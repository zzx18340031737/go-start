package main

import "fmt"

func main() {
	str := "Go语言"
	bytes := []byte(str)
	for i := 0; i < 2; i++ {
		bytes[i] = ' '
	}
	fmt.Println(string(bytes))
}