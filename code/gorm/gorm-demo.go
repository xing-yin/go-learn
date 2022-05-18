package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/17
 * @Desc: gorm 入门演示
 */

type Product struct {
	gorm.Model
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:price""`
}

// TableName maps to mysql table name
// 添加 TableName 方法，告诉 GORM 该 Models 映射到数据库中的哪张表
func (p *Product) TableName() string {
	return "product"
}

var (
	host     = pflag.StringP("host", "H", "127.0.0.1:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "123456", "Password for access to mysql, should be used pair with password")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// Parse command line flags
	// 使用 pflag 解析命令行的参数，通过命令行参数指定数据库的地址、用户名、密码和数据库名
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	// 创建完数据库连接后，会返回数据库实例 db ，之后就可以调用 db 实例中的方法，完成数据库的 CURD 操作。
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 1.Auto migration for given models ==> 不建议在正式的代码中自动迁移表结构
	db.AutoMigrate(&Product{})

	// 2.Insert the value into database
	if err := db.Create(&Product{Code: "D11", Price: 100}).Error; err != nil {
		log.Fatalf("Create error：%v", err)
	}
	PrintProducts(db)

	// 3.Find first record that match given conditions
	product := &Product{}
	if err := db.Where("code=?", "D11").First(&product).Error; err != nil {
		log.Fatalf("Get product error:%v", err)
	}

	// 4.Update value in database,if the value doesn't have primary key,will insert it
	product.Price = 200
	if err := db.Save(product).Error; err != nil {
		log.Fatalf("update product error: %v", err)
	}
	PrintProducts(db)

	// 5.Delete value match given conditions
	if err := db.Where("code = ?", "D11").Delete(&Product{}).Error; err != nil {
		log.Fatalf("Delete product error: %v", err)
	}
	PrintProducts(db)
}

// PrintProducts List all products
func PrintProducts(db *gorm.DB) {
	products := make([]*Product, 0)
	var count int64
	d := db.Where("code like ?", "%D%").Offset(0).Limit(2).Order("id desc").Find(&products).Offset(-1).Limit(-1).Count(&count)
	if d.Error != nil {
		log.Fatalf("List products error:%v", d.Error)
	}

	log.Printf("\ttotalcount:%d", count)
	for _, product := range products {
		log.Printf("\tcode: %s, price:%d\n", product.Code, product.Price)
	}
}
