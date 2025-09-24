package main

import (
	"GolangTask3/task"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:12345@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	/******************************* SQL语句练习 *******************************/
	// task.ExecTask1(db)

	/******************************* Sqlx入门 *******************************/
	// task.ExecTask2(db)

	/******************************* 进阶gorm *******************************/
	task.ExecTask3(db)
}
