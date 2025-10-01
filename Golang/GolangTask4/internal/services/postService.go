package services

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/repositories"
	"errors"
)

type PostService struct {
	postRepository *repositories.PostRepository
}

func NewPostService() *PostService {
	return &PostService{
		postRepository: &repositories.PostRepository{},
	}
}

func (p *PostService) AddPost(post *models.Post, userID any) error {
	post.UserID = userID.(uint)
	return p.postRepository.AddPost(post)
}

func (p *PostService) GetAllPost() (*[]models.Post, error) {
	return p.postRepository.GetAllPost()
}

func (p *PostService) GetUserAllPost(userIDParam any) (*[]models.Post, error) {
	userID := userIDParam.(uint)
	return p.postRepository.GetAllPostByUserID(userID)
}

func (p *PostService) GetUserPostInfo(id uint) (*models.Post, error) {
	return p.postRepository.FindByID(id)
}

func (p *PostService) UpdatePost(postReq *models.Post, id uint, userIDParam any) (*models.Post, error) {
	prePost, err := p.postRepository.FindByID(id)
	if err != nil {
		return nil, errors.New("notfound")
	}

	if userIDParam.(uint) != prePost.UserID {
		return nil, errors.New("unauthorized")
	}

	prePost.Title = postReq.Title
	prePost.Content = postReq.Content

	if err := p.postRepository.Update(prePost); err != nil {
		return nil, err
	}

	return p.postRepository.FindByID(prePost.ID)
}

func (p *PostService) DeletePost(id uint, userIDParam any) error {
	post, err := p.postRepository.FindByID(id)
	if err != nil {
		return errors.New("notfound")
	}

	if post.UserID != userIDParam {
		return errors.New("unauthorized")
	}

	return p.postRepository.Delete(post.ID)
}
