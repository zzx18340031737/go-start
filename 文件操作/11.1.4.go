package main

import (
	"fmt"
	"os"
)

/**
* 创建目录 赵芝兴
* @author
 */

func createDir(path string, dirName string) {
	dirPath := path + dirName
	//0777也可以os.ModePerm
	err := os.Mkdir(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
	}
	os.Chmod(dirPath, 0777)
	fmt.Println("Create Dir =>" + path + dirName)
}

func main() {
	createDir("/Users/zhaozhixing/Downloads/", "xkm")
	//createDir("/Users/zhaozhixing/Downloads/","test")

}
