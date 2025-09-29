package database

import (
	"GolangTask4/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDataBase() {
	// 数据库连接
	db, err := gorm.Open(mysql.Open("root:12345@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		return
	}
}
