package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Uid      int
	Username string
	Gender   bool
	Password string
	Created  string
}

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

	rows, err := db.Query("select * from `user` where username=?", "john")
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Uid, &user.Username, &user.Gender, &user.Password, &user.Created)
		checkErr(err)
		log.Println(user)
	}
}
