package repositories

import (
	"GolangTask4/internal/models"
	"GolangTask4/pkg/database"
)

type PostRepository struct {
}

func (p *PostRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := database.DB.Debug().Preload("User").Where("id = ?", id).First(&post).Error
	return &post, err
}

func (p *PostRepository) GetAllPost() (*[]models.Post, error) {
	var posts []models.Post
	err := database.DB.Debug().Preload("Comment").Order("created_at DESC").Find(&posts).Error
	return &posts, err
}

func (p *PostRepository) GetAllPostByUserID(userID uint) (*[]models.Post, error) {
	var posts *[]models.Post
	err := database.DB.Debug().Preload("Comment").Where("user_id = ?", userID).Find(posts).Error
	return posts, err
}

func (p *PostRepository) AddPost(post *models.Post) error {
	return database.DB.Debug().Create(&post).Error
}

func (p *PostRepository) Update(post *models.Post) error {
	return database.DB.Debug().Save(post).Error
}

func (p *PostRepository) Delete(id uint) error {
	return database.DB.Debug().Delete(&models.Post{}, id).Error
}
