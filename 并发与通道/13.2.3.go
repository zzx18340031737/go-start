package main

import (
	"fmt"
	"time"
)

func Task1() {
	for {
		fmt.Println(time.Now().Format("15:04:05"), "正在处理Task1的任务！")
		time.Sleep(time.Second * 3)
	}
}

func main() {
	go Task1()
	time.Sleep(time.Second * 100)
}
