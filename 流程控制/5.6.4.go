package main

import "fmt"

func pyramid(n int) {
	// n表示总层数
	for i := 1; i <= n; i++ {
		//打印*前先打印空格，规律为总层数-当前层数
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		//k表示每层打印多少*，规律为2 * i - 1
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	n := 9
	pyramid(n)
}
