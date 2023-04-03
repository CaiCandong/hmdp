package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/interfaces/controller"
)

// RegisterUploadRoutes 注册用户相关的路由
func RegisterUploadRoutes(r *gin.RouterGroup) {
	uploadCtrl := controller.NewUploadController()
	// 上传文件
	r.POST("/blog", uploadCtrl.Create)

	// 删除文件
	r.GET("/blog/delete", uploadCtrl.Delete)
}
