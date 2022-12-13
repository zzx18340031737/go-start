package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	n := runtime.GOMAXPROCS(1)
	fmt.Println("先前的CPU核数设置为：", n)
	last := time.Now()
	for i := 0; i < 100000; i++ {
		go func() {
			//耗时任务
			a := 999999 ^ 999999
			a = a + 1
		}()
	}
	now := time.Now()
	fmt.Println(now.Sub(last))
}
