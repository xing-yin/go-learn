package main

import "database/sql"

// recordStats 记录用户浏览产品信息
func recordStats(db *sql.DB, userID, productID int64) (err error) {
	// 开启事务
	// 操作 product 和 product_views 两张数据表
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// 更新 product 表
	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	// product_views 表插入一条新数据
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id,product_id) values (?,?)", userID, productID); err != nil {
		return
	}
	return
}

func main() {
	// 注意：测试的过程中并不需要真正的连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_table?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// userID 为2的用户浏览了 productID 为 5 的产品
	if err = recordStats(db, 2, 5); err != nil {
		panic(err)
	}
}
