package main

import (
	"fmt"
	"runtime"
	"time"
)

func Task1() {
	defer fmt.Println("task1 stop")
	fmt.Println("task1 start")
	fmt.Println("task1 work")
}

func Task2() {
	defer fmt.Println("task2 stop")
	fmt.Println("task2 start")
	runtime.Goexit() //效果和return一样
	fmt.Println("task2 work")
}

func main() {
	go Task1()
	go Task2()
	time.Sleep(time.Second * 5)
}
