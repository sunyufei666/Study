package handlers

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		postService: services.NewPostService(),
	}
}

// AddPost 创建文章
func (p *PostHandler) AddPost(c *gin.Context) {
	var post models.Post
	var err error
	if err = c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.postService.AddPost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"post": post})
}

// GetAllPost 获取所有文章
func (p *PostHandler) GetAllPost(c *gin.Context) {
	posts, err := p.postService.GetAllPost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"posts": posts})
}

// UpdatePost 更新文章
func (p *PostHandler) UpdatePost(c *gin.Context) {
	var post *models.Post
	var err error
	if err = c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err = p.postService.UpdatePost(post)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "post not found" {
			status = http.StatusNotFound
		} else if err.Error() == "unauthorized to update this post" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &post)
}

// DeletePost 删除文章
func (p *PostHandler) DeletePost(c *gin.Context) {
	var post *models.Post
	var err error
	if err = c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = p.postService.DeletePost(post); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "post not found" {
			status = http.StatusNotFound
		} else if err.Error() == "unauthorized to delete this post" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功！"})
}
