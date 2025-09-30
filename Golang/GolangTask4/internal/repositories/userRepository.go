package repositories

import (
	"GolangTask4/internal/models"
	"GolangTask4/pkg/database"
)

type UserRepository struct{}

func (u *UserRepository) CreateUser(user *models.User) error {
	//var user1 models.User
	//database.DB.Where("username = ?", user.Username).First(&user1)
	//fmt.Println(user1)
	return database.DB.Debug().Create(&user).Error
}

func (u *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (u *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
