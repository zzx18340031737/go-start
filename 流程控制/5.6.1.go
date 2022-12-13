package main

import "fmt"

func fibonacci(n int) (res int) {
	a := 1
	b := 1
	for i := 2; i < n; i++ {
		c := b
		b = a + b
		a = c
	}
	return b
}

func main() {
	n := 9 //求斐波那契数列第9项的值
	fmt.Printf("斐波那契数列第%d项的值为%d", n, fibonacci(n))
}
