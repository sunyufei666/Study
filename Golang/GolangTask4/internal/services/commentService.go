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

func (co *CommentService) GetAllCommentByPostID(postID uint) (*[]models.Comment, error) {
	return co.commentRepository.GetAllCommentByPostID(postID)
}

func (co *CommentService) AddComment(comment *models.Comment, postID uint, userIDParam any) error {
	comment.PostID = postID
	comment.UserID = userIDParam.(uint)
	return co.commentRepository.AddComment(comment)
}

func (co *CommentService) DeleteComment(id uint, userIDParam any) error {
	comment, err := co.commentRepository.FindCommentById(id)
	userID := userIDParam.(uint)
	if err != nil {
		return errors.New("notfound")
	}

	if userID != comment.UserID {
		return errors.New("unauthorized")
	}

	return co.commentRepository.DeleteComment(comment.ID)
}
