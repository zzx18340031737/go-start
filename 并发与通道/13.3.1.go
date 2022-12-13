package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
	}()
	ch <- "test"
	time.Sleep(time.Second)
}
