package main

import "fmt"

// 事务操作示例
// 演示了一个简单的事务操作，该事物操作能够确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。
func transaction() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}

	sqlStr1 := "Update user set age=32 where id=?"
	res1, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql1 failed,err:%v\n", err)
		return
	}
	affected1, err := res1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=42 where id=?"
	res2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affected2, err := res2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affected1, affected2)
	if affected1 == 1 && affected2 == 1 {
		fmt.Println("transaction commit success")
		tx.Commit() // 事务提交
	} else {
		fmt.Println("transaction rollback success")
		tx.Rollback()
	}

	fmt.Println("transaction execute success")
}
