package main

import (
	"fmt"
	"regexp"
)

func main() {
	targetString := "hello world"
	re := regexp.MustCompile(`(\w)+`)
	res := re.FindStringIndex(targetString)
	fmt.Println(res)
}
