package main

import "fmt"

func fibonacci(n int) (res int) {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

func main() {
	n := 6 //求斐波那契数列第5项的值
	fmt.Printf("斐波那契数列第%d项的值为%d", n, fibonacci(n))
}
