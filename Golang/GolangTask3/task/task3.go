package task

import (
	"gorm.io/gorm"
)

type User struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Name    string
	PostNum uint
	Posts   []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	Title         string
	UserID        uint
	CommentNum    uint
	CommentStatus string
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

func (p *Post) BeforeCreate(db *gorm.DB) (err error) {
	err = db.Debug().Model(&User{ID: p.UserID}).Update("post_num", gorm.Expr("post_num + ?", 1)).Error
	return
}

type Comment struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Content string
	PostID  uint
}

func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	var post Post
	db.Debug().Find(&post, c.PostID)
	post.CommentNum -= 1
	if post.CommentNum == 0 {
		post.CommentStatus = "无评论"
	}
	db.Debug().Save(&post)
	return
}

func ExecTask3(db *gorm.DB) {
	// Task1：模型定义
	// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）
	// 1. 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）
	// 2. 编写Go代码，使用Gorm创建这些模型对应的数据库表
	//createCommentTable(db)
	//createPostTable(db)
	//createUserTable(db)

	// Task2：关联查询
	// 1. 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	//var allPosts []Post
	//db.Debug().Where("user_id = ?", "1").Preload("Comments").Find(&allPosts)
	//fmt.Println(allPosts)

	// 2. 编写Go代码，使用Gorm查询评论数量最多的文章信息
	//var postID uint
	//var postResult Post
	//db.Debug().Model(&Comment{}).Select("post_id").Group("post_id").Order("COUNT(*) desc").Scan(&postID)
	//db.Debug().Model(&Post{ID: postID}).Preload("Comments").Find(&postResult)
	//fmt.Println(postResult)

	// Task3：钩子函数
	// 1. 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
	//db.Debug().Create(&Post{Title: "标题标题标题标题", UserID: 2})
	// 2. 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	db.Debug().Delete(&Comment{ID: 10, PostID: 6})
}

func createCommentTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Comment{})
	if err != nil {
		panic("评论表创建失败")
	}
}

func createPostTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Post{})
	if err != nil {
		panic("文章表创建失败")
	}
}

func createUserTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&User{})
	if err != nil {
		panic("用户表创建失败")
	}
}
