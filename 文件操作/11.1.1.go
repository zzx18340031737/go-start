package main

import (
	"fmt"
	"io/ioutil"
)

/*
* ReadDir
* @author 赵芝兴
 */

func main() {
	dir, err := ioutil.ReadDir("/Users/zhaozhixing/Documents/")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range dir {
		fmt.Println(file.Name())
	}
}
