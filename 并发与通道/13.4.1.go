package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		default:
			time.Sleep(time.Second)
		}
	}
}
