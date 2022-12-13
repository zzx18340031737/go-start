package main

import "fmt"

func main() {
	c1 := complex(1, 2)
	fmt.Println(c1, "实部为：", real(c1))
	fmt.Println(c1, "虚部为：", imag(c1))
}
