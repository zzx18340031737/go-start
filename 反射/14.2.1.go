package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	if typeOfA.Kind() == reflect.Int {
		fmt.Printf("变量a的kind是int")
	} else {
		fmt.Printf("变量a的Kind不是int")
	}
}
