package main

import (
	"log"
	"net"
)

func main() {
	//创建一个UDP地址
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	// 建立连接
	conn, err := net.DialUDP("udp", nil, udpaddr)
	if err != nil {
		log.Fatal("连接失败！", err)
	}
	defer conn.Close()
	log.Println("连接成功！")
}
