package task

import (
	"errors"

	"gorm.io/gorm"
)

type Students struct {
	Id    uint `gorm:"primarykey;autoIncrement"`
	Name  string
	Age   uint
	Grade string
}

type Accounts struct {
	gorm.Model
	ID      uint `gorm:"primarykey;autoIncrement`
	Balance float64
}

type Transactions struct {
	gorm.Model
	ID            uint `gorm:"primarykey;autoIncrement`
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
	FromAccount   Accounts `gorm:"foreignKey:FromAccountId";references:ID`
	ToAccount     Accounts `gorm:foreignKey:ToAccountId;references:ID`
}

func ExecTask1(db *gorm.DB) {
	// Task1. 基本CRUD操作
	// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、age （学生年龄，整数类型）、 grade （学生年级，字符串类型）
	// createStudentsTable(db)

	// 1. 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
	// db.Debug().Exec("insert into students(`name`,`age`,`grade`) VALUES (?,?,?)", "张三", 20, "三年级")

	// 2. 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
	// var students []Students
	// db.Debug().Raw("select * from students where age > ?", 18).Scan(&students)
	// fmt.Println(students)

	// 3. 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
	// db.Debug().Exec("update students set Grade = ? where Name = ?", "四年级", "张三")

	// 4. 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录
	// db.Debug().Exec("insert into students(`name`,`age`,`grade`) VALUES (?,?,?)", "李四", 12, "三年级")
	// db.Debug().Exec("delete from students where age < ?", 15)

	// Task2. 事务语句
	// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）
	// createAccountsTable(db)
	// createTransactionsTable(db)

	// 1. 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务
	// accountA := Accounts{ID: 1, Balance: 500.12}
	// accountB := Accounts{ID: 2, Balance: 354.99}
	// db.Debug().Create(&accountA)
	// db.Debug().Create(&accountB)
	db.Transaction(func(tx *gorm.DB) error {
		var accountA, accountB Accounts
		tx.Debug().First(&accountA, 1)
		tx.Debug().First(&accountB, 2)
		if accountA.Balance < 100 {
			return errors.New("账户余额不足")
		}
		tx.Debug().Model(&accountA).Update("Balance", accountA.Balance-100)
		tx.Debug().Model(&accountB).Update("Balance", accountB.Balance+100)
		if err := tx.Debug().Create(&Transactions{FromAccountId: 1, ToAccountId: 2, Amount: 100}).Error; err != nil {
			return err
		}
		return nil
	})
}

func createStudentsTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Students{})
	if err != nil {
		panic("学生表创建失败")
	}
}

func createAccountsTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Accounts{})
	if err != nil {
		panic("账户表创建失败")
	}
}

func createTransactionsTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Transactions{})
	if err != nil {
		panic("转账记录表创建失败")
	}
}
