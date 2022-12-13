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
	//创建UDP服务
	conn, err := net.ListenUDP("udp", udpaddr)
	checkErr(err)
	defer conn.Close()
	log.Println("UDP服务创建成功！")

	var buf = make([]byte, 1024)
	conn.Read(buf)
	log.Println(string(buf))

	_, raddr, err := conn.ReadFromUDP(buf)
	conn.Write([]byte("Hello Write\r\n"))
	conn.WriteToUDP([]byte("Hello WriteToUDP\r\n"), raddr)
}
