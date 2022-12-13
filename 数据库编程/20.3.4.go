package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:Wei208935@tcp(127.0.0.1:3306)/xzs")
	checkErr(err)
	defer db.Close()

	// 验证连接的可用性
	err = db.Ping()
	checkErr(err)
	log.Println("数据库连接成功！")
	//使用预编译方式，等效于
	//	rs,err := db.Exec("INSERT INTO `user`(username,gender,password,created) VALUES (?,?,?,?)","yangling",1,"123456",time.Now())
	stmt, err := db.Prepare("INSERT INTO `user`(username,gender,password,created) VALUES (?,?,?,?)")
	defer stmt.Close()
	rs, err := stmt.Exec("Liuqiaoyu", 1, "123456", time.Now())
	checkErr(err)
	rowCount, err := rs.RowsAffected()
	checkErr(err)
	log.Printf("插入了 %d 行", rowCount)
}
