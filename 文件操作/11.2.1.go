package main

import (
	"fmt"
	"os"
)

/**
* 打开文件
* @author 赵芝兴
 */

func main() {
	//以读写方式打开文件，如果不存在，则创建
	file, err := os.OpenFile("/Users/zhaozhixing/Downloads/xkm.txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	file.Close()
}
