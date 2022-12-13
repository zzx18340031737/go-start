package main

import (
	"fmt"
	"os"
)

/**
* 删除空目录
* @author 赵芝兴
 */

func deleteNotEmptyDir(dirPath string) {
	fmt.Println("Delete Dir =>" + dirPath)
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete Success")
	}

}

func main() {
	deleteNotEmptyDir("/Users/zhaozhixing/Downloads/xkm")
}
