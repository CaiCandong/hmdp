package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/controller"
	"hmdp/internal/app/middleware"
)

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(r *gin.RouterGroup) {
	userCtrl := controller.NewUserController()
	r.Use(middleware.CurrentUser())
	{
		// 短信登录&注册
		r.POST("/code", userCtrl.SendCode)

		// 注册用户登录接口
		r.POST("/login", userCtrl.Login)

		// 注册用户注册接口
		r.POST("/register", userCtrl.Register)

		r.Use(middleware.AuthRequired())
		{
			// 用户个人主页
			r.GET("/me", userCtrl.Me)

			// 注册用户注销接口
			r.DELETE("/logout", userCtrl.Logout)

			// 查看用户详情
			r.GET("/info/:id", userCtrl.Info)

			// 注册修改个人信息接口
			r.PUT("/:id", userCtrl.Update)

		}
	}
}
