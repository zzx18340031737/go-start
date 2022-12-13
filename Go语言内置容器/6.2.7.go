package main

import "fmt"

func main() {
	var student = [...]string{"Tom", "Ben", "Peter"}
	var student1 = student[0:1]
	fmt.Println("student数组：", student)
	fmt.Println("student1切片", student1)
	student1 = append(student1, "Danny") //对student切片的元素添加，会覆盖引用数组对应的元素
	fmt.Println("扩充Danny后的student1切片：", student1, ",切片长度为：", len(student1), "，切片容量为：", cap(student1))
	fmt.Println("扩充Danny后的student数组：", student)
}
