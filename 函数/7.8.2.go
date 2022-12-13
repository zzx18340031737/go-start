package main

import "fmt"

func passByReferencce(mapNumReference map[int]int) {
	fmt.Printf("passByReference函数中变量mapNumReference地址为：%p\n", mapNumReference)
	fmt.Printf("passByReference函数中变量mapNumReference所属指针地址为：%p\n", &mapNumReference)
	mapNumReference[1] = 100
}

func main() {
	mapNum := map[int]int{1: 10}
	fmt.Printf("main函数中变量mapNum地址为:%p\n", &mapNum)
	fmt.Printf("main函数中变量mapNum所属指针的地址为：%p\n", &mapNum)
	fmt.Printf("mapNum变量值为：%d\n", mapNum)
	passByReferencce(mapNum)
	fmt.Printf("mapNum变量值为：%d\n", mapNum)
}
