package main

import "fmt"

func Log(args ...interface{}) {
	for num, arg := range args {
		fmt.Printf("Index => %d,Value => %v\n",num,arg)
	}
}

func main(){
	s := make([]interface{},3)
	s[0] = 1
	s[1] = "abc"
	s[2] = struct {
		Num int
	}{1}

	//可变长参数
	fmt.Println("====将切片拆散====")
	Log(s...)
	fmt.Println("====直接传入切片====")
	Log(s)
}