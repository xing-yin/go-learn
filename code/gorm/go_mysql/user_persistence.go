package main

import "fmt"

type user struct {
	id   int
	name string
	age  int
}

// 插入数据示例
func insertRowDemo() {
	sqlStr := "insert into user(name,age) values(?,?)"
	res, err := db.Exec(sqlStr, "Bob", 20)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, id is %d.\n", id)
}

// 更新数据示例
func updateRowDemo() {
	sqlStr := "update user set age=? where id=?"
	res, err := db.Exec(sqlStr, 31, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", affected)
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id,name,age from user where id = ?"
	var u user
	// todo 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库连接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan faile,err:%v", err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id,name,age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query faile,err:%v", err)
		return
	}
	// todo 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan faile,err:%v", err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	}
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id=?"
	res, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := res.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
