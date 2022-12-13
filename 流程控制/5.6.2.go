package main

import "fmt"

func fibonacci(n int) (res int) {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

func main() {
	n := 5 //求斐波那契数列第5项的值
	fmt.Printf("斐波那契数列第%d项的值为%d", n, fibonacci(n))
}
