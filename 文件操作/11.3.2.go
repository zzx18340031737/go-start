package main

import (
	"encoding/json"
	"fmt"
)

/**
* 编码JSON-struct
* @auther 赵芝兴
 */

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Sex      bool     `json:sex`
	Birthday string   `json:"birthday"`
	Company  string   `json:"company,omitempty"`
	Language []string `json:"language"`
}

func main() {
	//定义一个结构体变量
	person := Person{"小王", 24, true, "1995-01-01", "芝兴有限公司", []string{"Go", "PHP", "Python"}}
	// result,err := json.Marshal(person)
	result, err := json.MarshalIndent(person, "", "         ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result = ", string(result))
}
