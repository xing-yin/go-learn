package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 原生的 mysql 初始化演示

// 定义一个全局对象 db
var db *sql.DB

func initDB() (err error) {
	// DNS: Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True&loc=Local"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，因为给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)

	// 尝试与数据库建立连接（检查 dsn 是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
