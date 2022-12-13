package main

import "fmt"

func main() {
	var student = []string{"Tom", "Ben", "Peter", "Danny"}
	student = append(student[0:1], student[2:]...)
	fmt.Println("student切片：", student)
	fmt.Println("student切片长度：", len(student))
	fmt.Println("student切片容量：", cap(student))
}
