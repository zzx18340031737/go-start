package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 3.0
	b := 2.0
	fmt.Println(a / b)
	fmt.Println("变量a的类型为:", reflect.TypeOf(a))
	fmt.Println("变量b的类型为:", reflect.TypeOf(b))

}
