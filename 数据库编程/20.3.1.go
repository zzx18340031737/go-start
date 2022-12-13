package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:Wei208935@tcp(127.0.0.1:3306)/xzs")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//验证连接的可用性
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	log.Fatal("数据库连接成功！")

}
