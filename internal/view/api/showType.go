package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/controller"
)

// RegisterShopTypeRoutes 注册用户相关的路由
func RegisterShopTypeRoutes(r *gin.RouterGroup) {
	shopTypeCtrl := controller.NewShowController()

	r.GET("/list", shopTypeCtrl.List)
}
