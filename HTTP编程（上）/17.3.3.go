package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"strconv"
)

func main() {
	// 初始化
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

	// 在OnResponse之后被调用
	c.OnHTML("#ip_list", func(e *colly.HTMLElement) {
		e.ForEach("tr > td:nth-of-type(2)", func(i int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
	})

	// 请求出错时，打印错误原因
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println(e.Error())
	})

	// 设置visit的线程数
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
	})

	// 爬取完成后，在OnHTML之后被调用
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// 获取前三页的ip地址
	for i := 1; i <= 3; i++ {
		c.Visit("https://www.xicidaili.com/nn/" + strconv.Itoa(i))
	}
}
