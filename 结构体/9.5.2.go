package main

import "fmt"

type User struct {
	name  string
	Email string
}

func (u *User) changeName() { //指针类型接收者
	u.name = "Tom"
}

func main() {
	u := User{"Peter", "go@go.com"} //创建指针类型结构体实例
	fmt.Println("name:", u.name, "Email:", u.Email)
	u.changeName()
	fmt.Println("name:", u.name, "Email:", u.Email)
}
