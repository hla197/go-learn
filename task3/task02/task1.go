package main

import (
	"fmt"
	"github.com/glebarez/sqlite" // ← 纯 Go sqlite 驱动
	"gorm.io/gorm"
)

/**
使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

type Employee struct {
	Name       string
	Department string
	Salary     float64
	gorm.Model
}

func initDB() *gorm.DB {
	// 连接 SQLite（文件不存在会自动创建）
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

func createData(db *gorm.DB) {
	employees := []Employee{
		{Name: "张三", Department: "技术部", Salary: 5000},
		{Name: "李四", Department: "销售部", Salary: 4000},
		{Name: "王五", Department: "技术部", Salary: 6000},
		{Name: "赵六", Department: "财务部", Salary: 3000},
		{Name: "孙七", Department: "技术部", Salary: 7000},
		{Name: "周八", Department: "销售部", Salary: 5000},
	}
	result := db.Create(employees)
	if result.Error != nil {
		panic(result.Error)
	}
}

func main() {
	db := initDB()

	db.AutoMigrate(&Employee{})

	// 清空数据 AllowGlobalUpdate: true → 允许全表删除, 危险行为, 这是软删除
	// db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Employee{})
	// 硬删除
	db.Debug().Unscoped().Delete(&Employee{}, "1=1")

	createData(db)
	type resultEmp struct {
		ID         uint
		Name       string
		Department string
		Salary     float64
	}

	fmt.Println("使用Sqlx查询 employees 表中所有部门为 \"技术部\" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中")
	var reslut []resultEmp
	db.Table("employees").Select("name", "id", "department", "salary").Where("department = ?", "技术部").Scan(&reslut)
	fmt.Println(reslut)

	fmt.Println("使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。")
	var reslutMax resultEmp
	db.Table("employees").Select("name", "id", "department", "salary").Order("salary desc").First(&reslutMax)
	fmt.Println(reslutMax)
}
