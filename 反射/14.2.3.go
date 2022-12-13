package main

import (
	"fmt"
	"reflect"
)

func checkValue(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Int {
		//方法一
		var v1 int = int(v.Int())
		//方法二
		var v2 int = v.Interface().(int)
		fmt.Println(v1, v2)
	}
}

func main() {
	var num int = 10
	valueOfNum := reflect.ValueOf(num)
	fmt.Println("valueOfNum")
	checkValue(valueOfNum)

	valueOfNumPtr := reflect.ValueOf(&num)
	fmt.Println("valueOfNumPtr")
	checkValue(valueOfNumPtr)
}
