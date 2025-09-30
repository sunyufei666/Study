package handlers

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService *services.LoginService
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{
		loginService: services.NewLoginService(),
	}
}

func (l *LoginHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := l.loginService.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功！"})
}

func (l *LoginHandler) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenStr, err := l.loginService.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenStr, "user": user})
}
