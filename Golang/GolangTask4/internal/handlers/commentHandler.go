package handlers

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		commentService: services.NewCommentService(),
	}
}

// AddComment 添加评论
func (co *CommentHandler) AddComment(c *gin.Context) {
	var comment models.Comment
	var err error
	if err = c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postID, _ := strconv.ParseInt(c.Param("postID"), 10, 32)
	userID, _ := c.Get("userID")
	err = co.commentService.AddComment(&comment, uint(postID), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"comment": comment})
}

// DeleteComment 删除评论
func (co *CommentHandler) DeleteComment(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	userID, _ := c.Get("userID")
	if err := co.commentService.DeleteComment(uint(id), userID); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "post not found" {
			status = http.StatusNotFound
		} else if err.Error() == "unauthorized to delete this post" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除评论成功！"})
}

// GetPostAllComment 获取某篇文章的所有评论
func (co *CommentHandler) GetPostAllComment(c *gin.Context) {
	postID, _ := strconv.ParseInt(c.Param("postID"), 10, 32)
	comments, err := co.commentService.GetAllCommentByPostID(uint(postID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &comments)
}
