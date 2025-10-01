package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts"`
}

func (User) TableName() string {
	return "users"
}
