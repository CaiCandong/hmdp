package controller

import "github.com/gin-gonic/gin"

func NewBlogCommentHandler() *BlogCommentHandler {
	return &BlogCommentHandler{}
}

type BlogCommentHandler struct {
}

func (b *BlogCommentHandler) Comment(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
