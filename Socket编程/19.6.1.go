package main

import (
	"flag"
	"log"
	"net"
	"time"
)

func scan(ip string, port string) {
	conn, err := net.DialTimeout("tcp", ip+":"+port, time.Second*2)
	if err != nil {
		log.Fatal("端口未开放！")
	}
	defer conn.Close()
	log.Fatal("端口开放！")

}

func main() {
	ip := flag.String("h", "127.0.0.1", "指定主机IP")
	port := flag.String("p", "8080", "指定扫描的端口")
	log.Println("扫描的端口为", *ip, *port)

	//扫描单个端口
	scan(*ip, *port)
}
