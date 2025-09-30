package services

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/repositories"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	userRepo *repositories.UserRepository
}

func NewLoginService() *LoginService {
	return &LoginService{
		userRepo: &repositories.UserRepository{},
	}
}

func (l *LoginService) Register(user models.User) (err error) {
	// 检查username是否已存在
	if _, err := l.userRepo.GetUserByUsername(user.Username); err == nil {
		return errors.New("username already exists")
	}

	// 检查email是否已存在
	if _, err := l.userRepo.GetUserByEmail(user.Email); err == nil {
		return errors.New("email already exists")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(hashedPassword)

	err = l.userRepo.CreateUser(&user)
	if err != nil {
		return
	}
	return nil
}

func (l *LoginService) Login(user models.User) (tokenStr string, err error) {
	// 查询数据库中的用户
	storedUser, err := l.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err = token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return
	}
	return tokenStr, nil
}

func (l *LoginService) ValidateToken(tokenStr string) (*models.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["id"].(float64))
		return l.userRepo.FindByID(userID)
	}

	return nil, errors.New("invalid token")
}
