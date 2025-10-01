package repositories

import (
	"GolangTask4/internal/models"
	"GolangTask4/pkg/database"
)

type CommentRepository struct{}

func (co *CommentRepository) AddComment(comment *models.Comment) error {
	return database.DB.Debug().Create(comment).Error
}

func (co *CommentRepository) GetAllCommentByPostID(postID uint) (*[]models.Comment, error) {
	var comments []models.Comment
	err := database.DB.Debug().Where("post_id = ?", postID).Find(&comments).Error
	return &comments, err
}

func (co *CommentRepository) FindCommentById(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := database.DB.Debug().First(&comment, id).Error
	return &comment, err
}

func (co *CommentRepository) DeleteComment(id uint) error {
	return database.DB.Debug().Delete(&models.Comment{}, id).Error
}
