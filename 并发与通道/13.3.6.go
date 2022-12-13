package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int) {
	for val := range in {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
