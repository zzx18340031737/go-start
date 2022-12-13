package main

import (
	"fmt"
	"reflect"
)

type Number int

type Person struct {
}

func checkType(t reflect.Type) {
	if t.Kind() == reflect.Ptr {
		fmt.Printf("变量的类型名称%v,指向的变量为：", t.Kind())
		t = t.Elem()
	}
	fmt.Printf("变量的类型名称 => %v,类型种类 => %v \n", t.Name, t.Kind())
}

func main() {
	var num Number = 10
	typeOfNum := reflect.TypeOf(num)
	fmt.Println("typeOfNum")
	checkType(typeOfNum)

	var person Person
	typeOfPerson := reflect.TypeOf(person)
	fmt.Println("typeOfPerson")
	checkType(typeOfPerson)

	typeOfPersonPtr := reflect.TypeOf(&person)
	fmt.Println("typeOfPersonPtr")
	checkType(typeOfPersonPtr)
}
