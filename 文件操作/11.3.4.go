package main

import (
	"encoding/json"
	"fmt"
)

/**
* 解码JSON-struct
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
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("person = %+v", person)
}
