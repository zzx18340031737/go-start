package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.shandonglvmengwl.com/", nil)
	if err != nil {
		fmt.Println(err)
	}
	//设置request的Header，具体可参考HTTP协议
	request.Header.Set("Accept", "text/html,application/xhtml+xml, application/xml;q=0.9, */*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7, *;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	request.Header.Set("Accept-Language", "zh-CN, zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	response, err := client.Do(request)
	fmt.Printf("%#v", response.Request.Header)
}
