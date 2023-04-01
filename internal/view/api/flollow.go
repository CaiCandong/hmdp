package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/controller"
)

// RegisterFollowerRoutes 注册用户相关的路由
func RegisterFollowerRoutes(r *gin.RouterGroup) {
	followCtrl := controller.NewFollowerController()
	// 订阅
	r.PUT("/:id/:isFollow", followCtrl.Follow)
	// 查看订阅状态
	r.GET("/or/not/:id", followCtrl.IsFollow)
	// 公共
	r.GET("/common/:id", followCtrl.Common)
}
