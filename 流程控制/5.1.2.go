package main

import (
	"fmt"
	"runtime"
)

func main() {
	if num := runtime.NumCPU(); num >= 1 {
		fmt.Println("程序使用的CPU核数为：", num)
	}
}
