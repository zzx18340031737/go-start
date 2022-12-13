package main

import "fmt"

func main() {
	var student = [...]string{"Tom", "Ben", "Peter"}
	var student1 = student[1:3]
	var student2 = student1[0:1]
	fmt.Println("studnet数组：", student[:])
	fmt.Println("student1切片：", student1[:])
	fmt.Println("student2切片：", student2[:])
	fmt.Println("student数组地址为", &student[1])   //取student[1]元素的值
	fmt.Println("student1切片地址为", &student1[0]) //取student[0]元素的地址
	fmt.Println("student2切片地址为", &student2[0]) //取student[0]元素的地址
	fmt.Println("student1切片长度为:", len(student1))
	fmt.Println("student1切片容量为:", cap(student1))
	fmt.Println("student2切片长度为:", len(student2))
	fmt.Println("student2切片容量为:", cap(student2))
}
