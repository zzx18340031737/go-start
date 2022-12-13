package main

import (
	"fmt"
	"runtime"
)

func main() {
	if _, _, line, ok := runtime.Caller(0); ok == true { //错误所在位置
		err := fmt.Errorf("***Line %d error***", line)
		fmt.Println(err.Error())
	}
}
