package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
		fmt.Println(f)
	}
}
