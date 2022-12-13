package main

import (
	"fmt"
	"regexp"
)

func main() {
	targetString := "hello world"
	re := regexp.MustCompile(`o`)
	res := re.ReplaceAllString(targetString, "O")
	fmt.Println(res)
}
