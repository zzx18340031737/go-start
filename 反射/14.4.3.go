package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	age  int `json:"age"`
	string
}

func main() {
	person := Person{"小张", 10, "备注"}
	valueOfPerson := reflect.ValueOf(&person)
	typeOfPerson := reflect.TypeOf(&person)

	for i := 0; i < valueOfPerson.Elem().NumField(); i++ {
		fieldValue := valueOfPerson.Elem().Field(i)
		fieldType := typeOfPerson.Elem().Field(i)
		fmt.Printf("类型名：%v 可以寻址：%v 可以设置：%v \n", fieldType.Name, fieldValue.CanAddr(), fieldValue.CanSet())
	}

	fmt.Println("修改前：", person)
	// 必须满足可寻址和可导出两个条件才能修改变量值
	valueOfPerson.Elem().Field(0).SetString("xiao zhang")

	fmt.Println("修改后：", person)
}
