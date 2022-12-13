package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("len(ch) = %v,cap(ch) = %v\n", len(ch), cap(ch))
			ch <- i
		}
		close(ch)
	}()
	for data := range ch {
		fmt.Println(data)
	}
}
