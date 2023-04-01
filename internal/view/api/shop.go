package api

import "github.com/gin-gonic/gin"

// RegisterShopRoutes 注册用户相关的路由
func RegisterShopRoutes(r *gin.RouterGroup) {
	//userCtrl := controller.NewUserController()
	// 发布博客
	r.POST("")
	// 点赞
	r.PUT("/like/:id")
	// 获取与用户相关的博客
	r.GET("/of/me")
	// 获取点赞数
	r.GET("/like/:id")
	// 当前热榜
	r.GET("/hot")
	r.GET("/:id")
	r.GET("/of/follow")

}
