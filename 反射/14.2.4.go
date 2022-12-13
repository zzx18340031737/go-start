package main

import (
	"fmt"
	"reflect"
)

func Equal(a, b int) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	// 反射调用函数需使用ValueOf
	valueOfFunc := reflect.ValueOf(Equal)

	//构造函数参数
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}

	//通过反射调用函数计算

	result := valueOfFunc.Call(args)

	fmt.Println("函数运行结果：", result[0].Bool())
}
