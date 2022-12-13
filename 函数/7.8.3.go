package main

import "fmt"

func main() {
	//将4，5，6追加到sourceSlice切片中
	sourceSlice := []int{1, 2, 3}
	sourceSlice = append(sourceSlice, 4, 5, 6)
	fmt.Println("sourceSlice:", sourceSlice)

	//将sourceSlice切片中的元素复制到targetSlice切片中
	targetSlice := make([]int, 3)
	num := copy(targetSlice, sourceSlice)
	fmt.Println("复制成功的元素个数：", num)
	fmt.Println("targetSlice:", targetSlice)
	fmt.Println("targetSlice切片长度为:", len(targetSlice))
	fmt.Println("targetSlice切片容量为:", len(targetSlice))
}
