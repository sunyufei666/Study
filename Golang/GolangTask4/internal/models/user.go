package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email"`
}

func (User) TableName() string {
	return "users"
}
