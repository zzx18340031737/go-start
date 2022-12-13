package main

import (
	"encoding/json"
	"fmt"
)

/**
* 编码JSON-map
* @auther 赵芝兴
 */

func main() {
	//创建一个map
	m := make(map[string]interface{}, 6)
	m["name"] = "小王"
	m["age"] = 24
	m["sex"] = true
	m["birthday"] = "1995-01-01"
	m["company"] = "芝兴有限公司"
	m["language"] = []string{"Go", "PHP", "Python"}
	//编写成JSON
	result, _ := json.Marshal(m)
	resultFormat, _ := json.MarshalIndent(m, "", "   ")
	fmt.Println("result = ", string(result))
	fmt.Println("resultFormat = ", string(resultFormat))
}
