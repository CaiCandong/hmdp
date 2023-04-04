package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/services"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/interfaces/controller"
)

// RegisterShopRoutes 注册用商户相关的路由
func RegisterShopRoutes(r *gin.RouterGroup) {
	shopHandler := controller.ShopHandler{
		ShopService: &services.ShopService{
			ShopRepo: &mysql.ShopRepo{DB: mysql.DB},
			ShopReq:  assembler.ShopReq{},
			ShopRsp:  assembler.ShopRsp{},
		},
	}
	// 发布博客
	r.POST("")
	// 点赞
	r.PUT("/like/:id")
	// 获取与用户相关的博客
	r.GET("/of/type", shopHandler.OfType)
	// 获取点赞数
	r.GET("/like/:id")
	// 当前热榜
	r.GET("/hot")
	r.GET("/:id")
	r.GET("/of/follow")

}
