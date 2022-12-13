package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
)

func main() {
	dbname := "db连接字符串"
	db, err := sql.Open("mysql", dbname) //ok
	//数据库操作
	_, _ = db, err
}
