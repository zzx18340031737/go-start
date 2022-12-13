package main

import (
	"fmt"
	"os"
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
	data := "我是赵芝兴"
	for i := 0; i < 3; i++ {
		//写入byte的slice数据
		file.Write([]byte(data))
	}
	file.Close()
}
