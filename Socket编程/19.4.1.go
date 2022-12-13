package main

import (
	"log"
	"net"
)

func main() {
	//尝试连接本地1234端口
	conn, err := net.Dial("udp", ":1234")
	if err != nil {
		log.Fatal("连接失败！", err)
	}
	defer conn.Close()
	log.Println("连接成功！")
}
