package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err := client.Do(request)
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	//打印Body
	fmt.Println(string(res))
}
