package main

import (
	"fmt"
	"os"
)

/**
* 读文件
* @auther 赵芝兴
 */

func ReadFile(path string) {
	//读取文件内容
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	//创建byte的slice用于接受文件读取数据
	buf := make([]byte, 1024)
	fmt.Println("以下是文件内容：")
	//循环读取
	for {
		//Read函数会改变文件当前偏移量
		len, _ := file.Read(buf)
		//读取字节数为0时跳出循环
		if len == 0 {
			break
		}
		fmt.Println(string(buf))
	}
	file.Close()
}

func main() {
	ReadFile("/Users/zhaozhixing/Downloads/xkm.txt")
}
