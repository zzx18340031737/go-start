package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
* WalkDir
* @author 赵芝兴
 */

//获取指定目录及所有子目录下的所有文件
func WalkDir(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.walk() returned %v/n", err)
	}
}

func main() {
	WalkDir("/Users/zhaozhixing/Downloads/")
}
