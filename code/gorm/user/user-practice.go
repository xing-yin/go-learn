package hook

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

/**
* @Author: Alan Yin
* @Date: 2022/5/17
* @Desc:
 */

type User struct {
	gorm.Model
	Name     string
	Age      uint8
	Birthday time.Time
}

type UserVO struct {
	ID   uint
	Name string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dns := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 大型项目中使用的连接池配置示例
	//sqlDB, err := db.DB()
	//sqlDB.SetMaxIdleConns(100)                 // 设置MySQL的最大空闲连接数（推荐100）
	//sqlDB.SetMaxOpenConns(100)                 // 设置MySQL的最大连接数（推荐100）
	//sqlDB.SetConnMaxLifetime(time.Second * 10) // 设置MySQL的空闲连接最大存活时间（推荐10s）

	db.AutoMigrate(&User{})

	// ------------------------------------------------------------------------------------
	// I.创建记录
	user := User{
		Name:     "Alan",
		Age:      20,
		Birthday: time.Now(),
	}
	reslut := db.Create(&user) // 根据数据指针来创建
	if reslut.Error != nil {
		log.Fatalf("Create error：%v", reslut.Error)
	}
	fmt.Println("user.ID:", user.ID)
	fmt.Println("reslut.RowsAffected:", reslut.RowsAffected)

	// batch insert:better performance
	var users = []User{
		{Name: "Bob", Age: 10},
		{Name: "Cindy", Age: 11},
	}
	db.Create(&users)
	// 等价于： db.CreateInBatches(&users, 2)

	for _, u := range users {
		fmt.Println(u.ID)
	}

	// ------------------------------------------------------------------------------------
	// II.删除记录
	// DELETE from users where id = 10 AND name = "jinzhu";
	reslut2 := db.Where("name=?", "Cindy").Delete(&User{})
	fmt.Println("reslut.RowsAffected:", reslut2.RowsAffected)

	// DELETE from users where id = 10;
	reslut3 := db.Delete(&User{}, 10)
	fmt.Println("reslut.RowsAffected:", reslut3.RowsAffected)

	reslut4 := db.Where("name in (?)", []string{"Bob", "Cindy"}).Delete(&User{})
	fmt.Println("reslut.RowsAffected:", reslut4.RowsAffected)

	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;
	reslut5 := db.Where("age=?", 11).Delete(&User{})
	fmt.Println("reslut5.RowsAffected:", reslut5.RowsAffected)

	user2 := &User{}
	// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
	reslut6 := db.Where("age=?", 20).Find(&user2)
	fmt.Println("reslut6.RowsAffected:", reslut6.RowsAffected)
	fmt.Println("user2:", user2)

	users2 := []User{}
	reslut7 := db.Unscoped().Where("age=10").Find(&users2)
	fmt.Println("reslut7.RowsAffected:", reslut7.RowsAffected)
	fmt.Println("users2:", users2)

	// 永久删除
	//reslut8 := db.Unscoped().Where("age=?", 10).Delete(&User{})
	//fmt.Println("reslut8.RowsAffected:", reslut8.RowsAffected)

	// ------------------------------------------------------------------------------------
	// III.更新记录
	user3 := &User{}
	reslut31 := db.Where("name=?", "Alan").First(&user3)
	if reslut31.Error != nil {
		log.Fatalf("find error：%v", reslut.Error)
	}
	user3.Birthday = time.Now()
	user3.Age = 21
	reslut32 := db.Save(&user3)
	if reslut32.Error != nil {
		log.Fatalf("update1 error：%v", reslut.Error)
	}

	// 指定更新单列
	// UPDATE users SET age=200, updated_at='2013-11-17 21:34:10' WHERE name='colin';
	db.Model(&User{}).Where("name=?", "Cindy").Update("age", 22)

	// 指定更新多列
	db.Model(&User{}).Where("name", "Cindy").Updates(User{Name: "hello", Age: 18})

	// ------------------------------------------------------------------------------------
	// IV.查询数据
	// a.检索单个记录
	user4 := &User{}
	reslut41 := db.First(&user4)
	fmt.Println("user4:", user4)
	fmt.Println("reslut7.RowsAffected:", reslut41.RowsAffected)

	user5 := &User{}
	reslut42 := db.Last(&user5)
	// 获取最后一条记录（主键降序）// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//reslut42 := db.First(&user5)
	if errors.Is(reslut42.Error, gorm.ErrRecordNotFound) {
		fmt.Println("reslut42：ErrRecordNotFound")
	} else {
		fmt.Println("user5:", user5)
		fmt.Println("reslut42.RowsAffected:", reslut42.RowsAffected)
	}

	// b.查询所有符合条件的记录
	users3 := make([]*User, 0)
	// SELECT * FROM users WHERE name <> 'Kar';
	db.Where("name <> ?", "Kar").Find(&user3)
	fmt.Println("users3:", users3)

	// c.智能选择字段
	uservo1 := UserVO{}
	// SELECT `id`, `name` FROM `users` LIMIT 10;
	reslut43 := db.Model(&User{}).Limit(10).Find(&uservo1)
	if reslut43.Error == nil {
		fmt.Println("uservo1:", uservo1)
	}

}
