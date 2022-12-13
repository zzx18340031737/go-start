package main

import (
	"log"
	"net"
)

func main() {
	//监听8080端口
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("服务器启动 失败！")
	}
	defer l.Close()
	log.Println("服务器启动成功")
}
