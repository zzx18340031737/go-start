package main

import (
	"strings"
	"fmt"
)

func main(){
	str := "Go语言"
	index := strings.Index(str,"语")
	fmt.Println(index)
	fmt.Println(str[index:])
}