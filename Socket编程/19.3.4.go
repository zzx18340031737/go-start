package main

import (
	"log"
	"net"
)

func main() {
	// ncat - 1kp 1234
	conn, err := net.Dial("tcp", "www.shandonglvmengwl.com:8080")
	if err != nil {
		log.Fatal("连接失败！")
	}
	defer conn.Close()
	log.Println("连接成功")
	//发送数据
	conn.Write([]byte("test\n"))
	//接受数据
	var buf = make([]byte, 10)
	conn.Read(buf)
	log.Println(buf)
}
