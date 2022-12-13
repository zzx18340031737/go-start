package main

import "fmt"

func main() {
	student := make([]int, 1, 1)
	for i := 0; i < 8; i++ {
		student = append(student, i)
		fmt.Println("当前切片长度：", len(student), "当前切片容量：", cap(student))
	}
}
