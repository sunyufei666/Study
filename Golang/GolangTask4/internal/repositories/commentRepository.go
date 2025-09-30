package repositories

import (
	"GolangTask4/internal/models"
	"GolangTask4/pkg/database"
)

type CommentRepository struct{}

func (co *CommentRepository) AddComment(comment *models.Comment) error {
	return database.DB.Debug().Create(comment).Error
}

func (co *CommentRepository) FindCommentById(id uint) (*models.Comment, error) {
	var comment *models.Comment
	err := database.DB.Debug().Preload("User").Where("id = ?", id).First(comment).Error
	return comment, err
}

func (co *CommentRepository) DeleteComment(id uint) error {
	return database.DB.Debug().Delete(&models.Comment{}, id).Error
}
