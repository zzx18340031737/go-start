package main

import (
	"fmt"
	"regexp"
)

func main() {
	targetString := "hello world"
	matchString := "hello"
	match, err := regexp.MatchString(matchString, targetString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(match)

}
