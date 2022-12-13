package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Result struct {
	Person []Person
}

type Person struct {
	Name      string
	Age       string
	Interests Interests
}

type Interests struct {
	Interests []string
}

func main() {
	var res Result
	content, err := ioutil.ReadFile("test.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = xml.Unmarshal(content, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("XML文件解析后内容为：")
	fmt.Println(res)
}
