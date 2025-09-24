package task

import (
	"fmt"

	"gorm.io/gorm"
)

type Employee struct {
	ID         uint `gorm:primaryKey;autoIncrement`
	Name       string
	Department string
	Salary     float64
}

func (Employee) TableName() string {
	return "employees"
}

type Book struct {
	ID     uint `gorm:primaryKey;autoIncrement`
	Title  string
	Author string
	Price  float64
}

func (Book) TableName() string {
	return "books"
}

func ExecTask2(db *gorm.DB) {
	// Task1：使用SQL扩展库进行查询
	// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary
	// createEmployeesTable(db)

	// 1. 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
	// db.Debug().Create(&Employee{Name: "张三", Department: "管理部门", Salary: 10000})
	// db.Debug().Create(&Employee{Name: "李四", Department: "采购部门", Salary: 8000})
	// db.Debug().Create(&Employee{Name: "王五", Department: "技术部门", Salary: 15000})
	// db.Debug().Create(&Employee{Name: "赵六", Department: "技术部门", Salary: 15000})
	// db.Debug().Create(&Employee{Name: "周七", Department: "技术部门", Salary: 16000})
	// var resultArr []Employee
	// db.Debug().Where("Department", "技术部门").Find(&resultArr)
	// fmt.Println(resultArr)

	// 2. 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
	// var result Employee
	// db.Debug().Order("salary desc").First(&result)
	// fmt.Println(result)

	// Task2：实现类型安全映射
	// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price
	// createBooksTable(db)
	// 1. 定义一个 Book 结构体，包含与 books 表对应的字段
	// book1 := Book {ID: 1, Title: "基督山伯爵", Author: "大仲马", Price: 100}
	// 2. 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
	// db.Debug().Create(&Book{Title: "基督山伯爵", Author: "大仲马", Price: 100})
	// db.Debug().Create(&Book{Title: "活着", Author: "余华", Price: 120})
	var bookResult []Book
	db.Debug().Where("price > ?", 50).Find(&bookResult)
	fmt.Println(bookResult)
}
func createEmployeesTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Employee{})
	if err != nil {
		panic("员工表创建失败")
	}
}

func createBooksTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Book{})
	if err != nil {
		panic("图书表创建失败")
	}
}
