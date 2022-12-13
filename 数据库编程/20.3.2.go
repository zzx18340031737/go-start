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

	// 验证连接的可用性
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	log.Println("数据库连接成功！")

	// 创建一张数据表
	createTable := "CREATE TABLE `xzs`.`xkm` (" +
		"`uid` INT(10) NOT NULL AUTO_INCREMENT," +
		"`username` VARCHAR(64) NULL DEFAULT 1," +
		"`gender` TINYINT(1) NULL DEFAULT NULL," +
		"`password` VARCHAR(64) NULL DEFAULT NULL," +
		"`created` DATE NULL DEFAULT NULL," +
		"PRIMARY KEY (`uid`)" +
		");"

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("数据表创建失败！", err)
	}
	log.Println("数据表创建成功！")
}
