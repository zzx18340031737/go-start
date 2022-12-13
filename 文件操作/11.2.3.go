package main

import (
	"fmt"
	"os"
)

/**
* 读文件
* @auther 赵芝兴
 */

func ReadFile2(path string) {
	//读取文件内容
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	//创建byte的slice用于接受文件读取数据
	buf := make([]byte, 1024)
	fmt.Println("以下是文件内容：")
	//读取偏移9字节的数据
	_, _ = file.ReadAt(buf, 9)
	fmt.Println(string(buf))
	_ = file.Close()
}

func main() {
	ReadFile2("/Users/zhaozhixing/Downloads/xkm.txt")
}
