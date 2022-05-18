package hook

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/18
 * @Desc: 原生 sql 练习
 */

type Result struct {
	ID   int
	Name string
	Age  uint8
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dns := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Result{})
	result := Result{
		Name: "Alan" + strconv.Itoa(rand.Int()),
		Age:  20,
	}
	result1 := db.Create(&result) // 根据数据指针来创建
	if result1.Error != nil {
		log.Fatalf("Create error：%v", result1.Error)
	}

	var result2 Result
	db.Raw("select id,name,age from result where id =?", 3).Scan(&result2)
	fmt.Println(result2)
}
