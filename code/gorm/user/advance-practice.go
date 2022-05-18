package hook

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/18
 * @Desc: gorm 高级查询演示
 */

type User2 struct {
	gorm.Model
	Name     string
	Age      uint8
	Birthday time.Time
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dns := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User2{})
	user := User2{
		Name:     "Alan" + strconv.Itoa(rand.Int()),
		Age:      20,
		Birthday: time.Now(),
	}
	reslut := db.Create(&user) // 根据数据指针来创建
	if reslut.Error != nil {
		log.Fatalf("Create error：%v", reslut.Error)
	}

	// ------------------------------------------------------------------------------------
	fmt.Println("1.指定检索记录时的排序方式--------------------------------------------")
	// 1.指定检索记录时的排序方式
	users1 := make([]*User2, 0)
	result1 := db.Order("age desc,name").Find(&users1)
	fmt.Println("result1.RowsAffected", result1.RowsAffected)
	for _, u1 := range users1 {
		fmt.Println(u1)
	}

	// ------------------------------------------------------------------------------------
	// 2.Limit & Offset
	fmt.Println("2.Limit & Offset--------------------------------------------")
	users2 := make([]User2, 0)
	db.Limit(10).Offset(5).Find(&users2)
	for _, u1 := range users2 {
		fmt.Println(u1)
	}

	// ------------------------------------------------------------------------------------
	// 3.Distinct
	fmt.Println("3.Distinct-------------------------------------------")
	users3 := make([]*User2, 0)
	db.Distinct("name", "age").Order("name, age desc").Find(&users3)
	for _, u := range users3 {
		fmt.Println(u)
	}

	// 4.Count
	var count int64
	db.Model(&User2{}).Where("name= ?", "Alan").Count(&count)
	fmt.Println("count is:", count)

	var user4 User2
	db.Raw("select id,name,age from user2 where id =?", 3).Scan(&user4)
	fmt.Println(user4)
}
