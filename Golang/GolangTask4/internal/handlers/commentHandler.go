package handlers

import (
	"GolangTask4/internal/models"
	"GolangTask4/internal/services"
	"net/http"

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
	var comment *models.Comment
	var err error
	if err = c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = co.commentService.AddComment(comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"comment": comment})
}

// DeleteComment 删除评论
func (co *CommentHandler) DeleteComment(c *gin.Context) {
	var comment *models.Comment
	var err error
	if err = c.ShouldBindJSON(comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = co.commentService.DeleteComment(comment); err != nil {
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

func (co *CommentHandler) getPostAllComment(c *gin.Context) {
	var commentReq *models.Comment
	var comments *[]models.Comment
	var err error
	if err = c.ShouldBindJSON(&commentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if comments, err = co.commentService.GetAllCommentByPostID(commentReq.PostID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &comments)
}
