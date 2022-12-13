package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	//使用NopCloser创建一个实现io.Closer接口的body
	body := ioutil.NopCloser(strings.NewReader("user=admin&pass=admin"))
	req, err := http.NewRequest("POST", "https://www.shandonglvmengwl.com/", body)
	if err != nil {
		fmt.Println(err)
	}
	//打印输出body
	req_body, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(req_body))

	//POST提交表单需要添加此Header头
	req.Header.Set("Content--Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	// 在client使用Do方法后，再次打印Body发现Body已经关闭，使用ReadAll读取会  导致错误
	req_body, err = ioutil.ReadAll(resp.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(req_body))
}
