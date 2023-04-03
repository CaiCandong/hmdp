package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/interfaces/controller"
)

// RegisterBlogCommentRoutes 注册用户相关的路由
func RegisterBlogCommentRoutes(r *gin.RouterGroup) {
	blogCommentCtrl := controller.NewBlogCommentController()
	//	博客评论
	r.POST("/blog-comments", blogCommentCtrl.Comment)

}
