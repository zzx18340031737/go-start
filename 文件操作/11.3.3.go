package main

/**
* 解码JSON-map
* @auther 赵芝兴
 */

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
{
         "name": "小王",
         "age": 24,
         "Sex": true,
         "birthday": "1995-01-01",
         "company": "芝兴有限公司",
         "language": [
                  "Go",
                  "PHP",
                  "Python"
         ]
}`
	//创建一个map
	m := make(map[string]interface{}, 6)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("m = ", m)
	//类型断言
	for key, value := range m {
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s]的值类型为string，value = %s\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为int，value = %f\n", key, data)
		case bool:
			fmt.Printf("map[%s]的值类型为bool，value = %t\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string，value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface{}，value = %v\n", key, data)
		}
	}
}
