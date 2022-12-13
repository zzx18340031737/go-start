package main

import "fmt"

func passByValue(numPara int) { //值传递函数
	fmt.Printf("passByValue函数中变量numPara地址为：%p\n", &numPara)
	numPara = 100
}

func passByReference(numPara *int) { //引用传递函数
	fmt.Printf("passByReference函数中指针变量numPara地址为:%p\n", &numPara)
	fmt.Printf("passByReference函数中指针变量numPara指向的地址为:%p\n", numPara)
	*numPara = 100
}

func main() {
	num := 1
	fmt.Printf("main函数中变量num地址为:%p\n", &num)
	passByValue(num)
	fmt.Printf("num变量值为：%d\n", num)
	passByReference(&num)
	fmt.Printf("num变量值为：%d\n", num)
}
