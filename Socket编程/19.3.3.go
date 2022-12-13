package main

import (
	"log"
	"net"
	"time"
)

func main() {
	//设置超时
	conn, err := net.DialTimeout("tcp", "www.baidu.com:81", time.Second*2)
	if err != nil {
		log.Fatal("连接失败", err)
	}
	defer conn.Close()
	log.Println("连接成功")
}
