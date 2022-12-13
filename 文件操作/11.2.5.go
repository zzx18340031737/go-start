package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
* 写文件
* @auther 赵芝兴
 */

func main() {
	file, err := os.Create("/Users/zhaozhixing/Downloads/xkm.txt")
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 3; i++ {
		//按指定偏移量写入数据
		ix := i * 64
		file.WriteAt([]byte("我是数据邢"+strconv.Itoa(i)+"\r\t"), int64(ix))
	}
	file.Close()
}
