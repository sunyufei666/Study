package services

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/repositories"
	"errors"
)

type CommentService struct {
	commentRepository *repositories.CommentRepository
}

func NewCommentService() *CommentService {
	return &CommentService{
		&repositories.CommentRepository{},
	}
}

func (co *CommentService) AddComment(comment *models.Comment) error {
	return co.commentRepository.AddComment(comment)
}

func (co *CommentService) DeleteComment(comment *models.Comment) error {
	comment1, err := co.commentRepository.FindCommentById(comment.ID)
	if err != nil {
		return errors.New("notfound")
	}

	if comment1.UserID != comment.UserID {
		return errors.New("unauthorized")
	}

	return co.commentRepository.DeleteComment(comment.ID)
}
