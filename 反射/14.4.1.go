package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 5

	fmt.Printf("type:%T \n", reflect.TypeOf(a))
	fmt.Printf("value:%T \n", reflect.ValueOf(a))
}
