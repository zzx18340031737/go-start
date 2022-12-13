package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := &struct {
		Code int
		Msg  string
	}{}
	jsonData := `{"code":200,"msg":""}`
	if err := json.Unmarshal([]byte(jsonData), data); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code:", data.Code)
	fmt.Println("msg:", data.Msg)
}
