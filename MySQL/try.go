package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:20050315yisheng@tcp(127.0.0.1:3306)/YiSheng")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("成功连接到数据库")

	// 执行查询
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// 检查查询过程中是否有错误
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
