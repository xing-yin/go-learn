package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)

	// 1.新增数据
	// db.Prepare 用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	// stmt.Exec()函数用来执行stmt准备好的SQL语句
	res, err := stmt.Exec("alan", "big data", "2020-01-01")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("id is :", id)

	// 2.更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("alan-update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("affect is :", affect)

	// 3.查询数据
	// db.Query()函数用来直接执行Sql返回Rows结果。
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid is :", uid)
		fmt.Println("username is :", username)
		fmt.Println("department is :", department)
		fmt.Println("created is :", created)
	}

	// 4.删除数据
	//stmt,err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res,err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect,err = res.RowsAffected()
	//checkErr(err)
	//fmt.Println("affect is:",affect)

	// 关闭连接
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
