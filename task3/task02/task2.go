package main

import (
	"fmt"
	"github.com/glebarez/sqlite" // ← 纯 Go sqlite 驱动
	"gorm.io/gorm"
)

/**
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     int `gorm:"primaryKey"`
	Title  string
	Author string
	Price  float64
}

func initDB() *gorm.DB {
	// 连接 SQLite（文件不存在会自动创建）
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

func InitData(db *gorm.DB) {
	books := []Book{
		{Title: "Go 语言基础", Author: "Go 语言", Price: 50},
		{Title: "Go 语言进阶", Author: "Go 语言", Price: 80},
		{Title: "Go 语言实战", Author: "Go 语言", Price: 60},
		{Title: "Go 语言微服务", Author: "Go 语言", Price: 90},
		{Title: "Go 语言分布式", Author: "Go 语言", Price: 70},
		{Title: "Go 语言区块链", Author: "Go 语言", Price: 100},
	}
	db.Create(&books)
}

func main() {
	db := initDB()
	db.AutoMigrate(&Book{})
	db.Delete(&Book{}, 1)
	InitData(db)

	fmt.Println("使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。")
	var books []Book
	db.Where("price > ?", 50).Find(&books)
	fmt.Println(books)
}
