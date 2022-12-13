package main

import (
	"log"
	"net"
	"time"
)

func main() {
	//监听8080端口
	l, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatal("服务器启动失败！", err)
	}
	defer l.Close()
	log.Println("服务器启动成功")
	//阻塞等待用户连接
	c, err := l.Accept()
	// 设置连接超时时间
	c.SetDeadline(time.Now().Add(time.Second))
	//设置读取超时时间
	c.SetReadDeadline(time.Now().Add(time.Second))
	//设置写入超时时间
	c.SetReadDeadline(time.Now().Add(time.Second))
}
