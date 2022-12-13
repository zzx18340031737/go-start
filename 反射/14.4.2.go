package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 5
	valueOfA := reflect.ValueOf(a)

	fmt.Println(valueOfA.Interface())
}
