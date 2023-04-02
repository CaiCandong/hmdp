package controller

import "github.com/gin-gonic/gin"

type BlogCommentController interface {
	Comment(c *gin.Context)
}

func NewBlogCommentController() BlogCommentController {
	return &BlogCommentControllerImp{}
}

type BlogCommentControllerImp struct {
}

func (b *BlogCommentControllerImp) Comment(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
