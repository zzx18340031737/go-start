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

func (p Person) GetName() {
	fmt.Println(p.Name)
}

func main() {
	person := Person{"小张", 10, "备注"}
	valueOfPerson := reflect.ValueOf(person)

	//根据名字获取方法
	f := valueOfPerson.MethodByName("GetName")

	// 执行结构体方法
	f.Call([]reflect.Value{})
}
