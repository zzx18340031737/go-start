package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int32 = 1234567891
	fmt.Println("变量a值为:", a, "变量类型为:", reflect.TypeOf(a))
	fmt.Println("转换变量a类型为int16,变量a值变为", int16(a), ",变量a类型变为:", reflect.TypeOf(int16(a)))
}
