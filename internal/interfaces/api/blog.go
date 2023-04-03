package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/interfaces/controller"
)

// RegisterBlogRoutes 注册用户相关的路由
func RegisterBlogRoutes(r *gin.RouterGroup) {
	blogCtrl := controller.NewBlogController()
	// 发布博客
	r.POST("")
	// 点赞
	r.PUT("/like/:id")
	// 获取与用户相关的博客
	r.GET("/of/me")
	// 获取点赞数
	r.GET("/like/:id")
	// 当前热榜
	r.GET("/hot", blogCtrl.Hot)
	r.GET("/:id", blogCtrl.Find)
	r.GET("/of/follow")
}
