package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

func main() {
	GoMap := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		go writeMap(GoMap, i, i)
		go readMap(GoMap, i)
	}
	fmt.Println("Done")
}

func readMap(Gomap map[int]int, key int) int {
	lock.Lock() //读map操作前先加锁
	m := Gomap[key]
	lock.Unlock() //读完map后解锁
	return m
}

func writeMap(Gomap map[int]int, key int, value int) {
	lock.Lock() //写完map操作前先加锁
	Gomap[key] = value
	lock.Unlock() //写完map后解锁
}
