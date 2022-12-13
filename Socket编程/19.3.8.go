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
	//监听8080端口
	l, err := net.Listen("tcp", ":8080")
	checkErr(err)
	defer l.Close()
	log.Println("服务器启动成功！")
	//阻塞等待用户连接
	c, err := l.Accept()
	checkErr(err)

	defer c.Close()
	//读取打印接受的信息
	var buf = make([]byte, 10)
	log.Println("start to read from conn")
	n, err := c.Read(buf)
	checkErr(err)
	log.Println("接受字节数：", n, "接受字内容为：", string(buf))

}
