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
	err := p.postRepository.AddPost(post)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostService) GetAllPost() (*[]models.Post, error) {
	posts, err := p.postRepository.GetAllPost()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) UpdatePost(post *models.Post) (*models.Post, error) {
	prePost, err := p.postRepository.FindByID(post.ID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if post.UserID != prePost.UserID {
		return nil, errors.New("unauthorized to update this post")
	}

	if err := p.postRepository.Update(post); err != nil {
		return nil, err
	}

	return p.postRepository.FindByID(post.ID)
}

func (p *PostService) DeletePost(post *models.Post) error {
	post1, err := p.postRepository.FindByID(post.ID)
	if err != nil {
		return errors.New("post not found")
	}

	if post.UserID != post1.UserID {
		return errors.New("unauthorized to delete this post")
	}

	return p.postRepository.Delete(post.ID)
}
