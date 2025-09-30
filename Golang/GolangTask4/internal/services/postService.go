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

func (p *PostService) AddPost(post *models.Post) error {
	return p.postRepository.AddPost(post)
}

func (p *PostService) GetAllPost() (*[]models.Post, error) {
	return p.postRepository.GetAllPost()
}

func (p *PostService) GetUserAllPost(userID uint) (*[]models.Post, error) {
	return p.postRepository.GetAllPostByUserID(userID)
}

func (p *PostService) GetUserPostInfo(id uint) (*models.Post, error) {
	return p.postRepository.FindByID(id)
}

func (p *PostService) UpdatePost(post *models.Post) (*models.Post, error) {
	prePost, err := p.postRepository.FindByID(post.ID)
	if err != nil {
		return nil, errors.New("notfound")
	}

	if post.UserID != prePost.UserID {
		return nil, errors.New("unauthorized")
	}

	if err := p.postRepository.Update(post); err != nil {
		return nil, err
	}

	return p.postRepository.FindByID(post.ID)
}

func (p *PostService) DeletePost(post *models.Post) error {
	post1, err := p.postRepository.FindByID(post.ID)
	if err != nil {
		return errors.New("notfound")
	}

	if post.UserID != post1.UserID {
		return errors.New("unauthorized")
	}

	return p.postRepository.Delete(post.ID)
}
