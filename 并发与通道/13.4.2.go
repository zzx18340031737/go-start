package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case val := <-ch:
				fmt.Println(val)
			case <-time.After(time.Second * 3):
				fmt.Println("已超时3秒")
				done <- true
			}
		}
	}()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	<-done
	fmt.Println("程序终止")
}
