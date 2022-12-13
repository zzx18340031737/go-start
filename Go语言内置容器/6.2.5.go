package main

import "fmt"

func main() {
	var student []int
	student = make([]int, 2, 10)
	fmt.Println("student切片：", student)
	fmt.Println("student切片长度为:", len(student))
	fmt.Println("student切片容量为:", cap(student))
	fmt.Println("判定student切片是否为空：", student == nil)
}
