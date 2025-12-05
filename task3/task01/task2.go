package main

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite" // ← 纯 Go sqlite 驱动
	"gorm.io/gorm"
)

/**
 accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Account struct {
	ID      uint
	Balance int
	Name    string
}

type Transaction struct {
	ID            uint
	FromAccountID uint
	ToAccountID   uint
	Amount        int
}

func initDB() *gorm.DB {
	// 连接 SQLite（文件不存在会自动创建）
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	db := initDB()
	db.AutoMigrate(&Account{}, &Transaction{})
	a := Account{
		Balance: 0,
		Name:    "A",
	}
	db.Create(&a)
	b := Account{
		Balance: 1000,
		Name:    "B",
	}
	db.Create(&b)

	var as []Account
	db.Model(&Account{}).Find(&as)
	fmt.Println(as)

	var fromAccountId uint = a.ID
	var toAccountId uint = b.ID
	var Amount = 100

	err := db.Transaction(func(tx *gorm.DB) error {
		fmt.Println("开始转账")
		defer fmt.Println("结束转账")
		var fromAccount Account
		if err := tx.Where("id = ?", fromAccountId).First(&fromAccount).Error; err != nil {
			return errors.New("A 账户不存在")
		}
		if fromAccount.Balance < 100 {
			return errors.New("A 账户余额不足")
		}

		var toAccount Account
		if err := tx.Where("id = ?", toAccountId).First(&toAccount).Error; err != nil {
			return errors.New("B 账户不存在")
		}
		fromAccount.Balance -= Amount
		toAccount.Balance += 100

		if err := tx.Save(&toAccount).Error; err != nil {
			return errors.New("更新 A 账户信息失败")
		}
		if err := tx.Save(&fromAccount).Error; err != nil {
			return errors.New("更新 B 账户信息失败")
		}
		transaction := Transaction{
			FromAccountID: fromAccountId,
			ToAccountID:   toAccountId,
			Amount:        Amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return errors.New("创建转账信息失败")
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("转账成功")
	}

	fmt.Println("查询账户信息")
	db.Model(&Account{}).Find(&as)
	fmt.Println(as)
}
