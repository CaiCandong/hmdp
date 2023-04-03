package controller

import "github.com/gin-gonic/gin"

func NewBlogCommentController() *BlogCommentController {
	return &BlogCommentController{}
}

type BlogCommentController struct {
}

func (b *BlogCommentController) Comment(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
