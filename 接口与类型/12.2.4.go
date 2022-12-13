package main

import (
	"errors"
	"fmt"
)

//MySQL数据库操作

type Mysql struct{
	DBName string
	isConnect bool
}

// Redis数据库操作
type Redis struct {
	DBName string
}

func (mysql *Mysql)Connect() error{
	fmt.Println("Mysql Connect DB =>" +mysql.DBName)

	//do Connect
	mysql.isConnect = true
	if mysql.isConnect {
		fmt.Println("Mysql Connect Success!")
		return nil
	}else{
		return errors.New("Connect failure!")
	}
}

func (mysql *Mysql)Disconnect() error{
	//do Disconnect
	fmt.Println("Mysql Disconnect Success!")
	return nil
}

func (redis *Redis)Connect() error{
	fmt.Println("Redis Connect DB =>" +redis.DBName)
	//do Connect
	fmt.Println("Redis Connect Success!")
	return nil
}

func (redis *Redis)Disconnect() error{
	//do DisConnect
	fmt.Println("Redis DisConnect Success!")
	return nil
}

func main(){
	var mysql = Mysql{DBName:"Student"}
	fmt.Println("开始连接")
	mysql.Connect()
	//do something
	fmt.Println("断开连接")
	mysql.Disconnect()

	var redis = Redis{DBName:"teacher"}
	fmt.Println("开始连接")
	redis.Connect()
	//do something
	fmt.Println("断开连接")
	redis.Disconnect()
}
