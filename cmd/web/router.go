package main

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/middleware"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/view/api"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.EnableCookieSession())
	router.Use(middleware.CurrentUser(mysql.NewUserRepo()))

	// 绑定用户相关的路由
	api.RegisterUserRoutes(router.Group("/user"))
	// 绑定商铺类型相关的路由
	api.RegisterShopTypeRoutes(router.Group("/shop-type"))
	// 绑定博客相关的路由
	api.RegisterBlogRoutes(router.Group("/blog"))
	//blog := router.Group("blog")
	//{
	//	blog.GET("hot")
	//}

	return router
}
