package database

import (
	"GolangTask4/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() {
	// 数据库连接
	var err error
	DB, err = gorm.Open(mysql.Open("root:12345@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		return
	}
}
