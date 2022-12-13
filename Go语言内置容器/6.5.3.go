package main

import (
	"fmt"
	"sync"
)

func main() {
	var GoMap sync.Map
	for i := 0; i < 1000000; i++ {
		go writeMap(GoMap, i, i)
		go readMap(GoMap, i)
	}
	fmt.Println("Done")
}

func readMap(Gomap sync.Map, key int) int {
	res, ok := Gomap.Load(key) //线程安全读取
	if ok == true {
		return res.(int)
	} else {
		return 0
	}
}

func writeMap(GoMap sync.Map, key int, value int) {
	GoMap.Store(key, value) //线程安全设置
}
