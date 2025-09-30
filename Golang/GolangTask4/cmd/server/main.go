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
	other := r.Group("/api")
	other.Use(middleware.LoginAuthMiddleware(services.NewLoginService()))
	{

	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
