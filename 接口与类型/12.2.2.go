package main

import (
	"fmt"
	"errors"
)

// Mysql数据库操作
type Mysql struct {
	DBName string
	isConnect bool
}

func (mysql *Mysql)Connect() error {
	fmt.Println("Mysql Connect DB => "+ mysql.DBName)

	//do Connect
	mysql.isConnect = true

	if mysql.isConnect {
		fmt.Println("Mysql Connect Success!")
		return nil
	} else {
		return errors.New("Connect failure! ")
	}
}

func (mysql *Mysql)Disconnect() error {
	//do Disconnect
	fmt.Println("Mysql Disconnect Success!")
	return nil
}