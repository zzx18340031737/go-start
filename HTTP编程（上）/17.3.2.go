package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// NewCollector(Options ...func(*Collector)) *Collector
	//声明初始化NewCollecter对象时可以指定Agent，连接递归深度，URL过滤以及domain限制等。
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36"))

	// 发出请求时附的回调
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
		//r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
		fmt.Println("Visiting", r.URL)

	})

	// 打印响应状态码
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", r.StatusCode)
	})

	// 对响应的HTML元素处理
	c.OnHTML("title", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		fmt.Println("title:", e.Text)
	})

	c.Visit("https://www.baidu.com")
}
