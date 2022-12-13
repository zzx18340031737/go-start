package main

import (
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//创建一个UDP地址
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:1234")
	checkErr(err)
	//建立连接
	conn, err := net.DialUDP("udp", nil, udpaddr)
	checkErr(err)
	defer conn.Close()
	log.Println("连接成功！")
	//发送数据
	conn.Write([]byte("Hello\r\n"))
	//接受数据
	var buf = make([]byte, 1024)
	conn.Read(buf)
}
