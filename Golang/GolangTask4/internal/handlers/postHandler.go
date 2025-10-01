package handlers

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/services"
	"net/http"
	"strconv"

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

	userID, _ := c.Get("userID")
	err = p.postService.AddPost(&post, userID)
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
	var postReq models.Post
	var postResp *models.Post
	var err error
	if err = c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, _ := c.Get("userID")
	postResp, err = p.postService.UpdatePost(&postReq, uint(id), userID)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "notfound" {
			status = http.StatusNotFound
		} else if err.Error() == "unauthorized" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postResp)
}

// DeletePost 删除文章
func (p *PostHandler) DeletePost(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	userID, _ := c.Get("userID")
	if err := p.postService.DeletePost(uint(id), userID); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "notfound" {
			status = http.StatusNotFound
		} else if err.Error() == "unauthorized" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除文章成功！"})
}

// GetUserAllPost 获取某用户所有的文章
func (p *PostHandler) GetUserAllPost(c *gin.Context) {
	userID, _ := c.Get("userID")
	posts, err := p.postService.GetUserAllPost(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetUserPostInfo 获取某篇文章的信息
func (p *PostHandler) GetUserPostInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	post, err := p.postService.GetUserPostInfo(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
