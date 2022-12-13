package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = "我是字符串"
	typeOfa := reflect.TypeOf(a)
	fmt.Println("变量a的类型为：" + typeOfa.Name())
	valueOfa := reflect.ValueOf(a)
	if typeOfa.Kind() == reflect.String {
		fmt.Println("变量a的值为：" + valueOfa.String())
	}
}
