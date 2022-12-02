package main

import "fmt"

func ReturnData() (int, int) {
	return 10, 20
}
func main() {
	a, _ := ReturnData()
	_, b := ReturnData()
	fmt.Println(a, b)
}
