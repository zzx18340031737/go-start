package main

import (
	"fmt"
	"os"
)

/**
* 删除空目录
* @author 赵芝兴
 */

func deleteEmptyDir(dirPath string) {
	fmt.Println("Delete Dir =>" + dirPath)
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete Success")
	}

}

func main() {
	deleteEmptyDir("/Users/zhaozhixing/Downloads/xkm")
}
