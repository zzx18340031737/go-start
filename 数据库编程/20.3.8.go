package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkErrWithTx(err error, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
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

	var password string

	tx, err := db.Begin()
	checkErr(err)

	// 查找John的密码，如果密码为123123就将密码改为111111，否则不执行任何操作
	err = tx.QueryRow("select password from `user` where username=?", "john").Scan(&password)
	checkErrWithTx(err, tx)
	if password == "123123" {
		rs, err := db.Exec("update `user` set password=? where username=?", "111116g", "john")
		checkErrWithTx(err, tx)
		rowCount, err := rs.RowsAffected()
		checkErrWithTx(err, tx)
		if rowCount > 0 {
			log.Println("密码更新成功！")
		}
	}
	tx.Commit()
	log.Println("事物处理完成！")
}
