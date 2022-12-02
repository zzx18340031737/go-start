package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	a := "012345"
	b := "6789"
	var c bytes.Buffer //声明变量c，类型为字节缓冲
	c.WriteString(a) //写入字符串变量a内容
	c.WriteString(b) //写入字符串变量b内容
	fmt.Println(c.String())
	fmt.Println(reflect.TypeOf(c))

}