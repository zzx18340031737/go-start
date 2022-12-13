package main

func main() {
	ch := make(chan int)
	<-ch //阻塞main goroutine，信道ch被锁
}
