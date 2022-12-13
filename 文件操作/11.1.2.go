package main

import (
	"fmt"
	"io/ioutil"
)

/**
* ListDir
* @author 赵芝兴
 */

//获取指定目录下的所有文件，不进入下一级目录搜索

func ListDir(dirPth string) error {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return err
	}
	for _, fi := range dir {
		if fi.IsDir() { //忽略目录
			fmt.Println("目录：" + fi.Name())
		} else {
			fmt.Println("文件：" + fi.Name())
		}
	}
	return nil
}

func main() {
	ListDir("/Users/zhaozhixing/Documents/")
}
