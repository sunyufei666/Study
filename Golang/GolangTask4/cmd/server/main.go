package main

import (
	"GolangTask4/internal/handlers"
	"GolangTask4/internal/middleware"
	"GolangTask4/internal/services"
	"GolangTask4/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDataBase()

	// 初始化路由
	r := gin.Default()

	// 登录相关
	loginHandler := handlers.NewLoginHandler()
	login := r.Group("/api")
	{
		login.POST("/register", loginHandler.Register)
		login.POST("/login", loginHandler.Login)
	}

	// 其他（需要权限验证）
	postHandler := handlers.NewPostHandler()
	commentHandler := handlers.NewCommentHandler()
	other := r.Group("/api")
	other.Use(middleware.LoginAuthMiddleware(services.NewLoginService()))
	{
		post := other.Group("/post")
		{
			post.POST("getAllPost", postHandler.GetAllPost)
			post.POST("addPost", postHandler.AddPost)
			post.POST("getUserAllPost", postHandler.GetUserAllPost)
			post.POST("getUserPostInfo/:id", postHandler.GetUserPostInfo)
			post.POST("updatePost/:id", postHandler.UpdatePost)
			post.POST("deletePost/:id", postHandler.DeletePost)
		}

		comment := other.Group("/comment")
		{
			comment.POST(":postID/addComment", commentHandler.AddComment)
			comment.POST(":postID/getAllComment", commentHandler.GetPostAllComment)
			comment.POST("deleteComment/:id", commentHandler.DeleteComment)
		}
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
