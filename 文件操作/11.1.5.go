package main

/**
* 创建目录
* @author 赵芝兴
 */

import (
	"fmt"
	"os"
)

func createDirAll(path string, dirName string) {
	dirPath := path + "//" + dirName
	fmt.Println("Create Dir =>" + dirPath)
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Create Success")
	}
	os.Chmod(dirPath, 0777)
}

func main() {
	createDirAll("/Users/zhaozhixing/Downloads/", "dir1//dir2//dir3")
}
