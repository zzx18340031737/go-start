package main

import (
	"fmt"
	"net"
)

func tcpSend() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err == nil {
		defer conn.Close()
		fmt.Println("remote address:", conn.RemoteAddr())
		fmt.Println("error:", err)
	}
}

func main() {
	tcpSend()
}
