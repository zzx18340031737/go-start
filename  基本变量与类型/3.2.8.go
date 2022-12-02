package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int16 = 97
	fmt.Println("变量a值为:", a, "变量类型为:", reflect.TypeOf(a))
	b := int32(a)
	fmt.Println("变量b值为:", b, "变量类型为:", reflect.TypeOf(b))
	fmt.Println("转换变量b类型为string:", string(b))

}
