package main

import (
	"GolangTask4/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDataBase()

	// 初始化路由
	r := gin.Default()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

//func Register(c *gin.Context) {
//	var user User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	// 加密密码
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
//		return
//	}
//	user.Password = string(hashedPassword)
//
//	if err := db.Create(&user).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
//}

//func Login(c *gin.Context) {
//	var user User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	var storedUser User
//	if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//		return
//	}
//
//	// 验证密码
//	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//		return
//	}
//
//	// 生成 JWT
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"id":       storedUser.ID,
//		"username": storedUser.Username,
//		"exp":      time.Now().Add(time.Hour * 24).Unix(),
//	})
//
//	tokenString, err := token.SignedString([]byte("your_secret_key"))
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
//		return
//	}
//	// 剩下的逻辑...
//}
