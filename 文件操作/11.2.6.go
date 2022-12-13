package main

import (
	"fmt"
	"os"
)

/**
* 删除文件
* @auther 赵芝兴
 */

func main() {
	//删除文件
	err := os.Remove("/Users/zhaozhixing/Downloads/xkm/xkm.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除成功！")
	}

	//删除指定path下的文件
	err = os.RemoveAll("/Users/zhaozhixing/Downloads/xkm")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除成功！")
	}
}
