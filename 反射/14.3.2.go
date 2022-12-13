package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int `json:"age"`
	string
}

func main() {
	person := Person{"小张", 10, "备注"}
	valueOfPerson := reflect.ValueOf(person)

	fmt.Printf("person的字段数量：%v\n", valueOfPerson.NumField())

	//通过下标访问获取字段值
	fmt.Println("Field")
	field := valueOfPerson.Field(1)
	fmt.Printf("字段值：%v \n", field.Int())

	// 通过字段名获取字段值
	field = valueOfPerson.FieldByName("Age")
	fmt.Println("FieldByName")
	fmt.Printf("字段值：%v \n", field.Interface())

	//通过下标索引获取字段值
	field = valueOfPerson.FieldByIndex([]int{0})
	fmt.Println("FieldByIndex")
	fmt.Printf("字段值：%v \n", field.Interface())
}
