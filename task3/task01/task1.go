package main

import (
	"fmt"
	"github.com/glebarez/sqlite" // ← 纯 Go sqlite 驱动
	"gorm.io/gorm"
)

/**
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

func initDB() *gorm.DB {
	// 连接 SQLite（文件不存在会自动创建）
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

type Student struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	Age   uint
	Grade string
}

func main() {
	db := initDB()

	db.AutoMigrate(&Student{})

	fmt.Println("-----------------------------")
	fmt.Println("向 students 表中插入一条新记录，学生姓名为 \"张三\"，年龄为 20，年级为 \"三年级\"")
	db.Debug().Create(&Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	})

	var student []Student
	fmt.Println("students 表中所有年龄大于 18 岁的学生信息。")
	db.Debug().Where("age > ?", 18).Find(&student)
	fmt.Println(student)

	fmt.Println("-----------------------------")
	fmt.Println("students 表中姓名为 \"张三\" 的学生年级更新为 \"四年级\"。")
	// students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//DB.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	db.Debug().Model(&Student{}).Where("name = ?", "张三").Updates(map[string]interface{}{"grade": "四年级", "age": 22})

	db.Model(&Student{}).Find(&student)
	fmt.Println(student)

	fmt.Println("-----------------------------")
	fmt.Println("删除 students 表中年龄小于 15 岁的学生记录")
	db.Debug().Where("age < ?", 15).Delete(&Student{})
	db.Debug().Model(&Student{}).Find(&student)
}
