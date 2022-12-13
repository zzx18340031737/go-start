package main

import "fmt"

func main() {
	fmt.Println("start now")
	defer fmt.Println("这是第一句defer语句")
	defer fmt.Println("这是第二句defer语句")
	defer fmt.Println("这是第三句defer语句")
	defer fmt.Println("这是第四句defer语句")
	fmt.Println("end")
}
