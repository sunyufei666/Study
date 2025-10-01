package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint
	User    User
	Comment []Comment
}

func (Post) TableName() string {
	return "posts"
}
