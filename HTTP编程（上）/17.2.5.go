package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http.Get实际上是DefaultClient.Get(url),Get函数是高度封装的，只有一个参数url
	//对于一般的http Request是可以使用，但是不能定制Request
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
